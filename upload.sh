#!/usr/bin/env bash
echo "This script builds and uploads 'check' binary to s3"
echo "Requires 'go' installed with linux_386 cross compilation and 'aws'"
echo ""

echo "Building..."

go-bindata -o ./gatherers.go gatherers/...
GOOS=linux GOARCH=386 CGO_ENABLED=0 go build -o koding-kernel
tar -cf koding-kernel.tar koding-kernel

# this isn't an authorization check; just to let others know this is dangerous
echo "This script will now upload gather binary to S3. THIS IS DANGEROUS OPERATION. ARE YOU SURE?! Enter your name to authorize:"
read name

if [[ -z $name ]]
then
  echo "Please enter your name...exiting"
  exit 1
fi

echo
echo "$name has authorized uploading..."
echo

echo "Uploading to s3..."
aws s3 cp koding-kernel.tar s3://koding-gather

echo "Cleaning up..."
rm koding-kernel koding-kernel.tar

echo "WARNING: You'll need to manually change permissions to this: http://take.ms/i9PpZ This'll be automated soon."
echo "Done"
