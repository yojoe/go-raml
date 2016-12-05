#!/bin/bash

go generate

# build
go install -v ./...
