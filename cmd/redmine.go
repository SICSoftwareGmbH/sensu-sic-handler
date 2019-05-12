// Copyright © 2019 SIC! Software GmbH

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// redmineCmd represents the "redmine" command
var redmineCmd = &cobra.Command{
	Use:   "redmine",
	Short: "Redmine related commands",
}

func init() {
	rootCmd.AddCommand(redmineCmd)

	redmineCmd.PersistentFlags().String(
		"redmine-url",
		"",
		"The redmine url to import data from")

	redmineCmd.PersistentFlags().String(
		"redmine-token",
		"",
		"The redmine token used for authentication")

	viper.BindPFlag("redmine-url", redmineCmd.PersistentFlags().Lookup("redmine-url"))
	viper.BindPFlag("redmine-token", redmineCmd.PersistentFlags().Lookup("redmine-token"))
}
