// Copyright © 2019 SIC! Software GmbH

package cmd

import (
	"os"
	"fmt"

	"github.com/go-redis/redis"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile     string
	redisClient *redis.Client
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
		"redis-host",
		"localhost",
		"The redis hostname")

	rootCmd.PersistentFlags().Int(
		"redis-port",
		6379,
		"The redis port")

	rootCmd.PersistentFlags().Int(
		"redis-db",
		0,
		"The redis db number")

	viper.BindPFlag("redis-host", rootCmd.PersistentFlags().Lookup("redis-host"))
	viper.BindPFlag("redis-port", rootCmd.PersistentFlags().Lookup("redis-port"))
	viper.BindPFlag("redis-db", rootCmd.PersistentFlags().Lookup("redis-db"))
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

	redisAddr := fmt.Sprintf("%s:%d", viper.GetString("redis-host"), viper.GetInt("redis-port"))
	redisClient = redis.NewClient(&redis.Options{Addr: redisAddr, DB: viper.GetInt("redis-db")})
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
