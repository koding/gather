#!/usr/bin/env bash
echo "This script builds and uploads 'check' binary to s3"
echo "Requires 'go' installed with linux_386 cross compilation and 'aws'"
echo ""

echo "Building..."

go-bindata -o ./gatherers.go gatherers/...
GOOS=linux GOARCH=386 CGO_ENABLED=0 go build
tar -cf gather.tar gather

echo "This script will now upload to S3. ARE YOU SURE?! Enter your name to proceed:"
read name

echo
echo "$name' has authorized uploading..."
echo

echo "Uploading to s3..."
aws s3 cp gather.tar s3://gather-vm

echo "Cleaning up..."
rm gather gather.tar

echo "Done"
