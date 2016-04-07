#!/bin/bash
set -ex
go get -u github.com/tools/godep
go get -u github.com/jteeuwen/go-bindata/...
go get -u github.com/Jumpscale/go-raml
cd $GOPATH/src/github.com/Jumpscale/go-raml
sh build.sh

