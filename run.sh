#!/bin/sh

GATHER=$1
if [ -z "$GATHER" ]
then
  GATHER="analytics"
fi

export GATHER=$GATHER

go-bindata -o ./gatherers.go gatherers/...
go build -o koding-kernel
./koding-kernel
rm koding-kernel
