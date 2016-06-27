#!/bin/bash

go-bindata -pkg templates -nometadata  -prefix codegen -o codegen/templates/bindata.go codegen/templates/*

# apidocs.zip
cd codegen/apidocs/html; rm -f ../apidocs_html.zip ; zip -rq  ../apidocs_html.zip *
cd -
go-bindata -nometadata -pkg apidocs -prefix codegen/apidocs -o codegen/apidocs/apidocs_html_zip.go codegen/apidocs/apidocs_html.zip

# date files
cd codegen/date
go-bindata -nometadata -pkg date -o bindate.go date_only.go  datetime.go datetime_only.go datetime_rfc2616.go  time_only.go
cd -

# build
go install -v ./...
