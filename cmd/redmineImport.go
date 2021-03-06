// Copyright © 2019 SIC! Software GmbH

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	etcd "go.etcd.io/etcd/clientv3"

	"github.com/SICSoftwareGmbH/sensu-sic-handler/redmine"
)

// redmineImportCmd represents the "redmine import" command
var redmineImportCmd = &cobra.Command{
	Use:   "import",
	Short: "Import recipients from redmine",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {
			terminateWithError("invalid argument(s) received")
			return
		}

		if viper.GetString("redmine-url") == "" {
			terminateWithError("redmine url is empty")
			return
		}

		if viper.GetString("redmine-token") == "" {
			terminateWithError("redmine token is empty")
			return
		}

		etcdClient, err := etcd.New(etcdConfig)
		if err != nil {
			terminateWithError("unable to connect to etcd")
			return
		}
		defer etcdClient.Close()

		err = redmine.Import(viper.GetString("redmine-url"), viper.GetString("redmine-token"), etcdClient)
		if err != nil {
			terminateWithError(err.Error())
			return
		}
	},
}

func init() {
	redmineCmd.AddCommand(redmineImportCmd)
}
