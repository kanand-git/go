/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// to generate the cli with viper support
// cobra-cli init --viper

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dock",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// we will ask user to provide the name of the file,
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "path to the JSON/Yaml config file")
	err := rootCmd.MarkPersistentFlagRequired("config")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
	if err != nil {
		fmt.Println("Error while binding config file", err)
		return
	}
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		fmt.Println("No config file provided")
		return
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	//ReadInConfig would read the config from the file, and then we can use the config in our project
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
