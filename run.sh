#!/bin/sh

GATHER=$1
if [[ -z "$GATHER" ]]
then
  GATHER="analytics"
fi

export GATHER=$GATHER

go-bindata -o ./gatherers.go gatherers/...
go build
./gather
rm gather
