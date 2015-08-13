#!/usr/bin/env bash
echo "This script builds and uploads 'check' binary to s3"
echo "Requires 'go' installed with linux_386 cross compilation and 'aws'"
echo ""

echo "Building..."

go-bindata -o ./gatherers.go gatherers/...
GOOS=linux GOARCH=386 CGO_ENABLED=0 go build
tar -cf gather.tar gather

# this isn't an authorization check; just to let others know this is dangerous
echo "This script will now upload gather binary to S3. THIS IS DANGEROUS OPERATION. ARE YOU SURE?! Enter your name to authorize:"
read name

echo
echo "$name has authorized uploading..."
echo

echo "Uploading to s3..."
aws s3 cp gather.tar s3://koding-gather

echo "Cleaning up..."
rm gather gather.tar

echo "Done"
