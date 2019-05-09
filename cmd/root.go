// Copyright © 2019 SIC! Software GmbH

package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/spf13/cobra"
)

var (
	redisHost   string
	redisPort   int
	redisDB     int
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

	rootCmd.PersistentFlags().StringVar(&redisHost,
		"redis-host",
		os.Getenv("REDIS_HOST"),
		"The redis hostname, defaults to value of REDIS_HOST env variable")

	rootCmd.PersistentFlags().IntVar(&redisPort,
		"redis-port",
		-1,
		"The redis port, defaults to value of REDIS_PORT env variable")

	rootCmd.PersistentFlags().IntVar(&redisDB,
		"redis-db",
		-1,
		"The redis db, defaults to value of REDIS_DB env variable")
}

// setup default values and variables
func initConfig() {
	if redisHost == "" {
		redisHost = "localhost"
	}

	if redisPort == -1 {
		redisPort = intValueFromEnvWithDefault("REDIS_PORT", 6379)
	}

	if redisDB == -1 {
		redisDB = intValueFromEnvWithDefault("REDIS_DB", 0)
	}

	redisClient = redis.NewClient(&redis.Options{Addr: fmt.Sprintf("%s:%d", redisHost, redisPort), DB: redisDB})
}

func intValueFromEnvWithDefault(env string, defaultValue int) int {
	if os.Getenv(env) == "" {
		return defaultValue
	}

	i, err := strconv.Atoi(os.Getenv(env))
	if err != nil {
		return defaultValue
	}

	return i
}

func terminateWithHelpAndMessage(cmd *cobra.Command, msg string) {
	_ = cmd.Help()

	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, msg)
}
