#!/bin/bash

cd ./commands && go-bindata ./templates/... && sed -i -- 's/package\ main/package\ commands/g' bindata.go && rm -f bindata.go--
