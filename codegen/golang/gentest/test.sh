#!/bin/bash
go generate
pushd struct
	go build -v ./...
popd
rm -rf struct
