// Copyright © 2019 SIC! Software GmbH

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	redmineURL   string
	redmineToken string
)

// redmineCmd represents the "redmine" command
var redmineCmd = &cobra.Command{
	Use:   "redmine",
	Short: "Redmine related commands",
}

func init() {
	rootCmd.AddCommand(redmineCmd)

	redmineCmd.PersistentFlags().StringVar(&redmineURL,
		"redmine-url",
		os.Getenv("REDMINE_URL"),
		"The redmine url to import data from, defaults to value of REDMINE_URL env variable")

	redmineCmd.PersistentFlags().StringVar(&redmineToken,
		"redmine-token",
		os.Getenv("REDMINE_TOKEN"),
		"The redmine token used for authentication, defaults to value of REDMINE_TOKEN env variable")
}
