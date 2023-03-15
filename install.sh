#!/bin/bash

# Set the download URL, version, and binary name
VERSION="1.0.0"
REPO="abyssmemes/esv"
BINARY_NAME="esv"
if [ "$(uname)" == "Darwin" ]; then
  OS="darwin"
else
  OS="linux"
fi
URL="https://github.com/${REPO}/releases/download/v${VERSION}/${BINARY_NAME}-${OS}-amd64"

# Download and install the binary
curl -sSL -o ${BINARY_NAME} ${URL}
chmod +x ${BINARY_NAME}

# Move the binary to /usr/local/bin
sudo mv ${BINARY_NAME} /usr/local/bin

echo "ESV CLI tool has been installed successfully!"