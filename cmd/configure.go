package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var ConfigureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure ESV CLI tool",
	Run:   configure,
}

func configure(cmd *cobra.Command, args []string) {
	reader := bufio.NewReader(os.Stdin)

	awsRegion := prompt("Enter your AWS Region (default: us-west-2): ", "us-west-2", reader)
	vaultAddr := prompt("Enter your Vault address: ", "", reader)
	vaultToken := prompt("Enter your Vault token: ", "", reader)
	vaultSshRoleName := prompt("Enter your Vault SSH role name: ", "", reader)
	imageID := prompt("Enter your AWS EC2 Image ID (default: ami-0c94855ba95b798c7): ", "ami-0c94855ba95b798c7", reader)
	keyName := prompt("Enter your AWS EC2 Key Pair name: ", "", reader)

	viper.Set("awsRegion", awsRegion)
	viper.Set("vaultAddr", vaultAddr)
	viper.Set("vaultToken", vaultToken)
	viper.Set("vaultSshRoleName", vaultSshRoleName)
	viper.Set("imageID", imageID)
	viper.Set("keyName", keyName)

	home, err := homedir.Dir()
	if err != nil {
		fmt.Println("Error finding home directory:", err)
		os.Exit(1)
	}

	configPath := filepath.Join(home, ".esv")
	err = os.MkdirAll(configPath, 0700)
	if err != nil {
		fmt.Println("Error creating configuration directory:", err)
		os.Exit(1)
	}

	configFile := filepath.Join(configPath, "config.yml")
	err = viper.WriteConfigAs(configFile)
	if err != nil {
		fmt.Println("Error writing configuration file:", err)
		os.Exit(1)
	}

	fmt.Println("Configuration saved.")
}

func prompt(message, defaultValue string, reader *bufio.Reader) string {
	fmt.Print(message)

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}

	input = strings.TrimSpace(input)
	if input == "" {
		return defaultValue
	}

	return input
}