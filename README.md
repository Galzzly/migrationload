# migrationload

Migration Load is a simple tool that allows the user to perform back to back migrations a WANdisco Live Data Migrator (LDM) environment. 

## How to run
Download the binary to your source LDM node and run 

```
$ migrationload -p /path/to/migrate -m migration-name
```

This will load data on the source cluster into `/path/to/migrate`, then create a new scan only migration called `migration-name` which will start once successfully created.

This will continue to run until told to stop, and will output to stdout in the terminal.

## Options
There are a number of options that can be passed to `migrationload`:

|Long Operator| Short Operator | Default Value | Description |
|-------------|----------------|---------------|-------------|
|`-path` |`-p`| | The parent path to load the source data into.|
|`-migration`|`-m`| | The name to use for the migration.|
|`-num`|`-n`|1|The number of migrations to run side by side.|
|`-wait`|`-w`|1m|The time to wait between polling for migration completion.|

## Flow of operations
An overview of how the tool works:

1. The arguments passed to the tool will be parsed, and the internal configuration set appropriately.
2. Set additional configuration items that are gathered from the system:
    * A client to connect to the local HDFS filesystem
    * The client to connect to the LDM API
    * The target filesystem to migrate data to
        >**NOTE**: Currently, there is an assumption that there will only be a single filesystem
3. Start the migration process
    1. Create the parent path on the source filesystem
    2. Generate random data in the parent path
    3. Create the migration with the Scan Only Migration and Auto Start options set
    4. After the wait time has passed, the tool will poll the LDM API to check whether the migration has completed or not
        * If it has not completed, the tool will wait again
        * If it has completed, the tool will:
            * Delete the path on the target filesystem
            * Remove the migration via the LDM API
            * Repeat all of the steps to recreate the migration

## Internal Default Options
### Random File Generator
When the random data is generated, the following options are used:

|Option|Value|Description|
|------|-----|-----------|
|filesize|10000000|The maximum file size in bytes|
|depth|3|How deep the directory tree should go|
|width|2|The number of subdirectories per directory|
|files|15|The number of files|
|randomfanout|true|Whether to randomise the number of files per directory|
|concurrent|5|How many files to make at a time|

### Create Migration
When the migration is created via the LDM API, the following options are used:

|Option|Value|Description|
|------|-----|-----------|
|ScanOnlyMigration|true|Whether the migration is created to scan only, and not move on to by live afterwards|
|AutoStart|true|Whether the migration should start automatically when the creation is successful|

### Delete Request
When the LDM API call is made to delete the path on the target filesystem, the following option is used"

|Option|Value|Description|
|------|-----|-----------|
|Recursive|true|Whether to recursively delete the path|

## Limitations
* At present the tool will work only with a HDFS source filesystems. 
* If there are multiple target filesystems, it will use the first one that is returned by the LDM API.

## To Do
Will be looking to work on the following:

* The use of a configuration file to store the configuration properties rather than using flags
* Allow user to choose which target to use when there are multiple target filesystems available
* Additional source filesystems beyond HDFS
* Add a subcommand `setup` to make sure that any configuration items are in place.
    * Include adding a service setup for `systemctl` so that it can be ran in the background
    * Include logrotate setup, so that it logs in to `/var/log`