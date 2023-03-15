package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	vaultApi "github.com/hashicorp/vault/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var ConnectCmd = &cobra.Command{
	Use:   "connect [instance_id]",
	Short: "Connect to an AWS EC2 instance using Vault SSH engine",
	Args:  cobra.ExactArgs(1),
	Run:   connect,
}

func connect(cmd *cobra.Command, args []string) {
	instanceID := args[0]

	awsRegion := viper.GetString("awsRegion")
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(awsRegion),
	})
	if err != nil {
		log.Fatal("Error creating AWS session")
	}

	ec2Svc := ec2.New(sess)

	params := &ec2.DescribeInstancesInput{
		InstanceIds: []*string{aws.String(instanceID)},
	}
	resp, err := ec2Svc.DescribeInstances(params)
	if err != nil {
		log.Fatal("Error describing EC2 instance")
	}

	publicIP := *resp.Reservations[0].Instances[0].PublicIpAddress

	vaultAddr := viper.GetString("vaultAddr")
	vaultToken := viper.GetString("vaultToken")
	vaultSshRoleName := viper.GetString("vaultSshRoleName")
	client, err := vaultApi.NewClient(&vaultApi.Config{
		Address: vaultAddr,
	})
	if err != nil {
		log.Fatal("Error creating Vault client")
	}

	client.SetToken(vaultToken)
	sshSecrets, err := client.Logical().Write(fmt.Sprintf("ssh/creds/%s", vaultSshRoleName), map[string]interface{}{
		"username": "ec2-user",
		"ip":       publicIP,
	})
	if err != nil {
		log.Fatal("Error retrieving SSH credentials from Vault")
	}

	privateKey := sshSecrets.Data["private_key"].(string)

	keyPath := fmt.Sprintf("%s.pem", instanceID)
	err = os.WriteFile(keyPath, []byte(privateKey), 0600)
	if err != nil {
		log.Fatal("Error writing private key to file")
	}

	defer os.Remove(keyPath)

	sshCmd := exec.Command("ssh", "-i", keyPath, fmt.Sprintf("ec2-user@%s", publicIP))
	sshCmd.Stdin = os.Stdin
	sshCmd.Stdout = os.Stdout
	sshCmd.Stderr = os.Stderr

	fmt.Printf("Connecting to EC2 instance %s at %s...\n", instanceID, publicIP)
	err = sshCmd.Run()
	if err != nil {
		log.Fatal("Error connecting to EC2 instance via SSH")
	}
}