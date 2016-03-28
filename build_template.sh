#!/bin/bash

mkdir -p ./codegen/bindata && cd ./codegen/bindata && go-bindata ../templates/... && sed -i -- 's/package\ main/package\ bindata/g' bindata.go && rm -f bindata.go--
