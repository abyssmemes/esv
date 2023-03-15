package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ec2-vault-cli",
	Short: "A CLI tool to manage AWS EC2 instances and SSH connections using Vault",
}

func Execute() error {
	return rootCmd.Execute()
}