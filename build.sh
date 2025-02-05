#!/bin/bash

MODULE_PATH="main.go"

export GOOS=linux
export GOARCH=amd64
echo "Building for $GOOS/$GOARCH..."
go build -o ~/source/Publish/unzipper/unzipper $MODULE_PATH

export GOOS=windows
export GOARCH=amd64
echo "Building for $GOOS/$GOARCH..."
go build -o ~/source/Publish/unzipper/unzipper.exe $MODULE_PATH

echo "Done"