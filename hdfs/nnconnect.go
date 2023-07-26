package nnconnect

import (
	"fmt"
	"net"
	"os"
	"os/user"
	"strings"
	"time"

	"github.com/colinmarc/hdfs/v2"
	"github.com/colinmarc/hdfs/v2/hadoopconf"
	krb "github.com/jcmturner/gokrb5/v8/client"
	"github.com/jcmturner/gokrb5/v8/config"
	"github.com/jcmturner/gokrb5/v8/credentials"
)

func ConnectToNamenode() (*hdfs.Client, error) {
	return getClient()
}

func getClient() (*hdfs.Client, error) {
	namenode := os.Getenv("HADOOP_NAMENODE")
	conf, err := hadoopconf.LoadFromEnvironment()
	if err != nil {
		return nil, err
	}

	options := hdfs.ClientOptionsFromConf(conf)
	if namenode != "" {
		options.Addresses = []string{namenode}
	}

	if options.Addresses == nil {
		return nil, fmt.Errorf("cannot find namenode to connect to")
	}

	if options.KerberosClient != nil {
		options.KerberosClient, err = getKerberosClient()
		if err != nil {
			return nil, fmt.Errorf("problem with kerberos auth: %s", err)
		}
	} else {
		options.User = os.Getenv("HADOOP_USER_NAME")
		if options.User == "" {
			u, err := user.Current()
			if err != nil {
				return nil, fmt.Errorf("unable to determine user: %s", err)
			}

			options.User = u.Username
		}
	}

	dialFunc := (&net.Dialer{
		Timeout:   5 * time.Second,
		KeepAlive: 5 * time.Second,
		DualStack: true,
	}).DialContext

	options.NamenodeDialFunc = dialFunc
	options.DatanodeDialFunc = dialFunc

	client, err := hdfs.NewClient(options)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to namenode: %s", err)
	}

	return client, nil
}

func getKerberosClient() (*krb.Client, error) {
	conf := os.Getenv("KRB_CONFIG")
	if conf == "" {
		conf = "/etc/krb5.conf"
	}

	cfg, err := config.Load(conf)
	if err != nil {
		return nil, err
	}

	ccachepath := os.Getenv("KRBCCNAME")
	if strings.Contains(ccachepath, ":") {
		if strings.HasPrefix(ccachepath, "FILE:") {
			ccachepath = strings.SplitN(ccachepath, ":", 2)[1]
		} else {
			return nil, fmt.Errorf("unusable cache: %s", ccachepath)
		}
	} else if ccachepath == "" {
		u, err := user.Current()
		if err != nil {
			return nil, err
		}
		ccachepath = fmt.Sprintf("tmp/krbcc_%s", u.Uid)
	}

	ccache, err := credentials.LoadCCache(ccachepath)
	if err != nil {
		return nil, err
	}

	return krb.NewFromCCache(ccache, cfg)
}
