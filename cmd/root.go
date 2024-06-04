/*
Copyright © 2024 Ranarivola Herinavalona hernavalasco@gmail.com
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/elewis787/boa"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var masterPassword string

var rootCmd = &cobra.Command{
	Use:   "pass-manager",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

func Execute() {
	rootCmd.SetUsageFunc(boa.UsageFunc)
	rootCmd.SetHelpFunc(boa.HelpFunc)

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// define global flags accessible everywhere
	rootCmd.PersistentFlags().StringVarP(&masterPassword, "password", "p", "", "The master password for encryption mechanism")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		root, err := os.Getwd()
		cobra.CheckErr(err)

		// Search config in root directory with name "pass-manager" (without extension).
		viper.AddConfigPath(root)
		viper.SetConfigType("yaml")
		viper.SetConfigName("pass-manager")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
