#!/bin/bash

mkdir -p ./commands/bindata && cd ./commands/bindata && go-bindata ../templates/... && sed -i -- 's/package\ main/package\ bindata/g' bindata.go && rm -f bindata.go--
