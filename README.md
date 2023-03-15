# ESV CLI Tool

The ESV (EC2 SSH Vault) CLI tool is a command-line utility that simplifies the process of creating and connecting to AWS EC2 instances using HashiCorp Vault's SSH engine for secure access.

## Features

- Create AWS EC2 instances with a specified image ID and key pair name
- Connect to EC2 instances using Vault's SSH engine for secure access
- Easy configuration of AWS credentials, Vault address, token, and SSH role name

## Pre

You need to have configured AWS CLI on your machine ( https://github.com/aws/aws-cli )

## Installation

```bash
curl -sSL https://raw.githubusercontent.com/abyssmemes/esv/main/install.sh | bash
```

#or

Clone this repository to your local machine:

```bash
git clone https://github.com/abyssmemes/esv.git
```

Build the binary:

```bash
cd esv
go build -o esv
```

Move the binary to a directory in your $PATH:

```bash
sudo mv esv /usr/local/bin
```
## Configuration

Before using the ESV CLI tool, you need to configure it with your AWS and Vault credentials. Run the following command and provide the requested information:

```bash
esv configure
```

This will create a configuration file in ~/.esv/config.yml with your credentials.

## Usage

## Create an EC2 instance
To create an AWS EC2 instance, use the create command:

```bash
esv create
```

This will create an EC2 instance with the image ID and key pair name specified in the configuration.

## Connect to an EC2 instance
To connect to an EC2 instance using Vault's SSH engine, use the connect command followed by the instance ID:

```bash
esv connect <instance_id>
```

Replace <instance_id> with the ID of the instance you want to connect to. This command will obtain a signed SSH key from Vault and connect to the specified EC2 instance using the obtained key.