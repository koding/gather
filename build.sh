#!/usr/bin/env bash
go-bindata -o ./gatherers.go gatherers/
GOOS=linux GOARCH=386 CGO_ENABLED=0 go build
