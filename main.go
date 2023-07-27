package main

import (
	"context"
	crand "crypto/rand"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"path"
	"strconv"
	"sync"
	"syscall"
	"time"

	nnconnect "github.com/Galzzly/migrationload/hdfs"
	"github.com/Galzzly/migrationload/swagger"
	"github.com/antihax/optional"
	"github.com/colinmarc/hdfs/v2"
	flag "github.com/spf13/pflag"
)

const (
	waitTime     = 60 * time.Second
	renewTime    = 23 * time.Hour
	migName      = "TestMigration"
	FilenameSize = 16
)

type hdfsClient struct {
	*hdfs.Client
}

type options struct {
	filesize int64
	depth    int32
	files    int32
	width    int32

	randomfanout bool

	concurrent int32
}

var (
	defaultOpts = options{
		filesize:     10000000, // The maximum file size
		depth:        3,        // How deep the directory tree should go
		width:        2,        // The number of subdirectories per directory
		files:        15,       // The total number of files
		randomfanout: true,
		concurrent:   5,
	}
	alphabet      = []rune("abcdefghijklmnopqrstuvwxyz0123456789-_") // The characters available to use for filenames
	migrationOpts = &swagger.MigrationControllerApiNewMigrationWithIdOpts{
		ScanOnlyMigration: optional.NewBool(true), // Want to just perform a scan only migration
		AutoStart:         optional.NewBool(true), // Want the migration to start automatically on creation
	}
	deleteOpts = swagger.FileDeleteRequest{
		Recursive: true,
	}
)

type config struct {
	wait      time.Duration
	path      string
	migration string
	nummigs   int

	ldmapi           *swagger.APIClient
	targetfilesystem string
	hdfsclient       hdfsClient
	renew            time.Duration

	opts options
}

/*
Initialise the bare config set.
This will feed in from the flags that are provided to the command
e.g. migrationload -w 30s -p /testpath -m testmig -n 5
will create 5 migrations under paths /testpath0 to /testpath4, and
under testmig0 to testmig4
It will poll LDM every 30 seconds for each migration to check if it
has completed. If completed it will delete the migration and remove
the source directory to start over again.
*/
func (c *config) init(args []string) error {
	flags := flag.NewFlagSet(args[0], flag.ExitOnError)

	flags.DurationVarP(&c.wait, "wait", "w", waitTime, "Time to wait between checks")
	flags.StringVarP(&c.path, "path", "p", "", "Path of the base migration")
	flags.StringVarP(&c.migration, "migration", "m", "", "Name of the migration")
	flags.IntVarP(&c.nummigs, "num", "n", 1, "Number of migrations to run side by side")

	if err := flags.Parse(args[1:]); err != nil {
		return err
	}

	return nil
}

func init() {
	/*
		Check if one of the following environment variables are set
			HADOOP_HOME
			HADOOP_CONF_DIR

		If neither are set, then attempt to set HADOOP_HOME to /etc/hadoop
	*/
	if os.Getenv("HADOOP_HOME") == "" || os.Getenv("HADOOP_CONF_DIR") == "" {
		if e := os.Setenv("HADOOP_HOME", "/etc/hadoop"); e != nil {
			log.Panicf("Unable to set environment variable")
		}
	}
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	// Create an empty config set
	c := &config{}

	// Create the channel to catch any signals send to the tool
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	// Create a background function to monitor for the incoming signals in the channel
	// made above
	go func() {
		for {
			select {
			case s := <-signalChan:
				switch s {
				case syscall.SIGINT, syscall.SIGTERM:
					log.Printf("Got SIGINT/SIGTERM, exiting")
					cancel()
					os.Exit(1)
				case syscall.SIGHUP:
					log.Printf("Got SIGHUP, reloading...")
					c.init(os.Args)
				}
			case <-ctx.Done():
				log.Printf("Done!")
				os.Exit(1)
			}
		}
	}()

	defer func() {
		signal.Stop(signalChan)
		cancel()
	}()

	// Start the core of the tool, and wait for the returned error
	// If there is an error, print to Stderr & exit
	if err := run(ctx, c, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, c *config, out io.Writer) error {
	// Initialise the config, which will parse the flags provided
	c.init(os.Args)
	// Set the log output to the Writer provided
	log.SetOutput(out)

	// Get the HDFS client that will be used to build out the source paths
	// Add the client to the config set
	client, err := nnconnect.ConnectToNamenode()
	if err != nil {
		return err
	}
	c.hdfsclient = hdfsClient{client}

	c.renew = renewTime

	// Could do with doing a check to see if they are set via a config file
	// TODO: Set up config file.
	c.opts = defaultOpts

	// Create a new LDM client to communicate with the LDM API
	// Add the LDM client to the config set
	log.Printf("Getting LDM client...")
	cfg := &swagger.Configuration{
		BasePath:      "http://localhost:18080",
		DefaultHeader: make(map[string]string),
		UserAgent:     "MigrationLoadClient",
	}

	c.ldmapi = swagger.NewAPIClient(cfg)

	// Get the target filesystem.
	// For now, this will work on the premise that there will be only one source
	// and one target.
	r, targetfilesystems, httpresp, err := c.ldmapi.TargetFileSystemControllerApi.ListTargetFileSystems(ctx, &swagger.TargetFileSystemControllerApiListTargetFileSystemsOpts{Verbose: optional.NewBool(true)})
	if err != nil {
		log.Println("Requested:", r)
		log.Println(targetfilesystems)
		log.Println(httpresp)
		return fmt.Errorf("there is problem getting target filesystem: %v", err)
	}
	if httpresp.StatusCode != 200 {
		return fmt.Errorf("problem getting target filesystem: %s", swagger.HttpStatusMap[httpresp.StatusCode])
	}
	if len(targetfilesystems) == 0 {
		return fmt.Errorf("no target filesystems found")
	}
	c.targetfilesystem = targetfilesystems[0].FileSystemId

	// Now that we have all of the details that we need, we are able to break out and
	// start creating the migrations per the number requested by the user.
	for i := 0; i < c.nummigs; i++ {
		go c.migrate(ctx, out, i)
	}

	// Continuous loop that will renew the HDFS client every 23 hours
	// Typically only relevant for kerberos enabled clusters so that the
	// keytab is renewed
	for {
		select {
		case <-time.Tick(c.renew):
			client, err := nnconnect.ConnectToNamenode()
			if err != nil {
				log.Printf("Unable to renew client: %v", err)
				os.Exit(1)
			}
			c.hdfsclient = hdfsClient{client}
		}
	}
}

/*
The migrate function will perform the following operations
  - Initialise the root directory for the migration, the root
    directory will be <pathname>n where n is the number of the
    migration
  - Create the source data structure within the root directory
  - Set up the LDM migration with the name <migration_name>n
    where n is the number of the migration

-- The migration will be created as scan only, and will start

		automatically once created
	  - Poll for the period set (default 1 minute) to confirm if the
	    migration has completed successfully

-- If it has not completed, it will simply continue to wait

	until the next poll period

-- If it has completed successfully then it will remove clean

	up the current migration, removing the migrated files on
	both source and target, remove the migration, and then
	recreate the source path and structure before initiating a
	new migration
*/
func (c config) migrate(ctx context.Context, out io.Writer, num int) error {
	n := strconv.Itoa(num)
	migpath := c.path + n
	migname := c.migration + n
	log.Printf("%s: Initialising path... %s", migname, migpath)
	log.Printf("%s: Creating root directory...", migname)
	if err := c.hdfsclient.MkdirAll(migpath, 0755); err != nil {
		log.Printf("%s: Failed to create root directory %s: %v", migname, migpath, err)
		return fmt.Errorf("%s: problem creating root directory %s: %v", migname, migpath, err)
	}

	log.Printf("%s: Creating structure...", migname)
	if err := c.hdfsclient.WriteRandomFiles(migpath, 1, &c.opts); err != nil {
		return fmt.Errorf("%s: problem creating random files: %v", migname, err)
	}

	log.Printf("%s: Structure created successfully for %s", migname, migpath)
	log.Printf("%s: Setting up LDM migration...", migname)

	// Start the scan & migration
	migration, httpresp, err := c.ldmapi.MigrationControllerApi.NewMigrationWithId(ctx, migname, migpath, c.targetfilesystem, migrationOpts)
	if err != nil {
		return fmt.Errorf("%s: there is a problem starting initial migration: %v", migname, err)
	}
	if httpresp.StatusCode != 201 {
		return fmt.Errorf("%s: problem starting initial migration: %s", migname, swagger.HttpStatusMap[httpresp.StatusCode])
	}
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-time.Tick(c.wait):
			log.Printf("%s: Checking status of migration...", migname)
			// Get the status of the migration
			migration, httpresp, err = c.ldmapi.MigrationControllerApi.FetchMigration(ctx, migration.MigrationId)
			if err != nil {
				return fmt.Errorf("%s: problem fetching migration: %v", migname, err)
			}
			if httpresp.StatusCode != 200 {
				return fmt.Errorf("%s: problem fetching migration: %s", migname, swagger.HttpStatusMap[httpresp.StatusCode])
			}

			if migration.State != nil {
				state := *migration.State
				// If the state is not "COMPLETED", then continue to wait
				if state != swagger.COMPLETED_MigrationState {
					continue
				}
			}
			// The state is "COMPLETED" so we can clean up the current migration,
			// and start again with a new one.
			// Delete path in target
			response, httpresp, err := c.ldmapi.FileSystemControllerApi.DeleteByPath(ctx, deleteOpts, migpath, c.targetfilesystem, "targets")
			if err != nil {
				return fmt.Errorf("%s: error deleting path on target: %v", migname, err)
			}
			if httpresp.StatusCode != 200 {
				return fmt.Errorf("%s: error deleting path on target: %s", migname, swagger.HttpStatusMap[httpresp.StatusCode])
			}
			if !response {
				return fmt.Errorf("%s: unable to delete path on target: %s", migname, migpath)
			}

			// Delete the migration
			httpresp, err = c.ldmapi.MigrationControllerApi.DeleteMigration(ctx, migration.MigrationId)
			if err != nil {
				return fmt.Errorf("%s: error deleting migration: %v", migname, err)
			}
			if httpresp.StatusCode != 200 {
				return fmt.Errorf("%s: error deleting migration: %s", migname, swagger.HttpStatusMap[httpresp.StatusCode])
			}

			// Delete existing source data
			log.Printf("%s: Deleting source path...", migname)
			if err := c.hdfsclient.RemoveAll(migpath); err != nil {
				return fmt.Errorf("%s: problem deleting source path: %v", migname, err)
			}

			// Recreate the source path again
			log.Printf("%s: Creating structure...", migname)
			if err := c.hdfsclient.WriteRandomFiles(migpath, 1, &c.opts); err != nil {
				return fmt.Errorf("%s: problem creating random files: %v", migname, err)
			}

			// Generate new source data
			log.Printf("%s: Recreating structure...", migname)
			if err := c.hdfsclient.WriteRandomFiles(migpath, 1, &c.opts); err != nil {
				return fmt.Errorf("%s: problem creating random files: %v", migname, err)
			}

			log.Printf("%s: Structure created successfully for %s", migname, migpath)
			log.Printf("%s: Setting up LDM migration...", migname)
			migration, httpresp, err = c.ldmapi.MigrationControllerApi.NewMigrationWithId(ctx, migration.MigrationId, migpath, c.targetfilesystem, migrationOpts)
			if err != nil {
				return fmt.Errorf("%s: problem starting initial migration: %v", migname, err)
			}
			if httpresp.StatusCode != 201 {
				return fmt.Errorf("%s: problem starting initial migration: %s", migname, swagger.HttpStatusMap[httpresp.StatusCode])
			}
		}
	}
}

/*
WriteRandomFiles
*/
func (c hdfsClient) WriteRandomFiles(root string, depth int32, opts *options) error {
	var wg sync.WaitGroup
	numFiles := opts.files
	if opts.randomfanout {
		numFiles = rand.Int31n(numFiles)
	}
	wg.Add(int(numFiles))

	workers := make(chan int, opts.concurrent)
	for i := int32(0); i < numFiles; i++ {
		go c.WriteRandomFile(root, opts, &wg, workers)
	}

	if depth+1 <= opts.depth {
		numDirs := opts.depth
		for i := int32(0); i < numDirs; i++ {
			if err := c.WriteRandomDir(root, depth+1, opts); err != nil {
				return err
			}
		}
	}
	return nil
}

func (c hdfsClient) WriteRandomFile(root string, opts *options, wg *sync.WaitGroup, worker chan int) error {
	defer wg.Done()
	worker <- 1
	size := rand.Intn(FilenameSize-4) + 4
	name := RandomFilename(size)

	filepath := path.Join(root, name)
	f, err := c.Create(filepath)
	if err != nil {
		<-worker
		return err
	}

	if _, e := io.CopyN(f, crand.Reader, opts.filesize); e != nil {
		return e
	}
	<-worker
	return f.Close()
}

func (c hdfsClient) WriteRandomDir(root string, depth int32, opts *options) error {
	if depth > opts.depth {
		return nil
	}

	size := rand.Intn(FilenameSize-4) + 4
	name := RandomFilename(size)
	root = path.Join(root, name)
	if err := c.MkdirAll(root, 0755); err != nil {
		return err
	}

	return c.WriteRandomFiles(root, depth, opts)
}

func RandomFilename(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(b)
}
