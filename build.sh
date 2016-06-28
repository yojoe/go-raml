#!/bin/bash

go-bindata -pkg templates -nometadata  -prefix codegen -o codegen/templates/bindata.go codegen/templates/*

# date files
cd codegen/date
go-bindata -nometadata -pkg date -o bindate.go date_only.go  datetime.go datetime_only.go datetime_rfc2616.go  time_only.go
cd -

# build
go install -v ./...
