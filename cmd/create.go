package cmd

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an AWS EC2 instance",
	Run:   create,
}

func create(cmd *cobra.Command, args []string) {
	awsRegion := viper.GetString("awsRegion")
	imageID := viper.GetString("imageID")
	keyName := viper.GetString("keyName")

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(awsRegion),
	})

	if err != nil {
		log.Fatal("Error creating AWS session")
	}

	ec2Svc := ec2.New(sess)

	instanceInput := &ec2.RunInstancesInput{
		ImageId:      aws.String(imageID),
		InstanceType: aws.String("t2.micro"),
		KeyName:      aws.String(keyName),
		MinCount:     aws.Int64(1),
		MaxCount:     aws.Int64(1),
	}

	result, err := ec2Svc.RunInstances(instanceInput)
	if err != nil {
		log.Fatal("Error creating EC2 instance")
	}

	instanceID := *result.Instances[0].InstanceId
	fmt.Printf("EC2 instance created with ID: %s\n", instanceID)
}