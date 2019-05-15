// Copyright © 2019 SIC! Software GmbH

package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	etcd "go.etcd.io/etcd/clientv3"
)

var (
	cfgFile    string
	etcdConfig etcd.Config
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sensu-sic-handler",
	Short: "The Sensu Go SIC! Software handler",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile,
		"config",
		"c",
		os.Getenv("SIC_CONFIG"),
		"Configuration file to use, defaults to value of SIC_CONFIG env variable")

	rootCmd.PersistentFlags().String(
		"etcd-endpoints",
		"localhost:2379",
		"Endpoints for etcd, defaults to localhost:2379")

	_ = viper.BindPFlag("etcd-endpoints", rootCmd.PersistentFlags().Lookup("etcd-endpoints"))
}

// setup default values and variables
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)

		// If a config file is found, read it in.
		err := viper.ReadInConfig()
		if err != nil {
			terminateWithError(err)
		}
	}

	etcdConfig = etcd.Config{
		Endpoints:   strings.Split(viper.GetString("etcd-endpoints"), ","),
		DialTimeout: 2 * time.Second,
	}
}

func terminateWithHelpAndMessage(cmd *cobra.Command, msg string) {
	_ = cmd.Help()

	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, msg)
}

func terminateWithError(err error) {
	fmt.Println(err)
	os.Exit(1)
}
