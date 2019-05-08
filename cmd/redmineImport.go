// Copyright © 2019 SIC! Software GmbH

package cmd

import (
	"github.com/spf13/cobra"

	"sensu-sic-handler/redmine"
)

// redmineImportCmd represents the "redmine import" command
var redmineImportCmd = &cobra.Command{
	Use:   "import",
	Short: "Import recipients from redmine",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {
			terminateWithHelpAndMessage(cmd, "invalid argument(s) received")
			return
		}

		if redmineURL == "" {
			terminateWithHelpAndMessage(cmd, "redmine url is empty")
			return
		}

		if redmineToken == "" {
			terminateWithHelpAndMessage(cmd, "redmine token is empty")
			return
		}

		err := redmine.Import(redmineURL, redmineToken, redisClient)
		if err != nil {
			terminateWithHelpAndMessage(cmd, err.Error())
			return
		}
	},
}

func init() {
	redmineCmd.AddCommand(redmineImportCmd)
}
