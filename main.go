package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/abyssmemes/esv/cmd"

	"github.com/mitchellh/go-homedir"
)

func main() {
	cobra.OnInitialize(loadConfig)

	rootCmd := &cobra.Command{
		Use:   "esv",
		Short: "A CLI tool to create and connect to AWS EC2 instances using Vault SSH engine",
	}

	rootCmd.AddCommand(cmd.ConfigureCmd)
	rootCmd.AddCommand(cmd.CreateCmd)
	rootCmd.AddCommand(cmd.ConnectCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func loadConfig() {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	configPath := filepath.Join(home, ".esv")
	viper.AddConfigPath(configPath)
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	viper.SetDefault("awsRegion", "us-west-2")
	viper.SetDefault("imageID", "ami-0c94855ba95b798c7")

	err = viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}