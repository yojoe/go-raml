#!/bin/bash

./build_template.sh

# apidocs.zip
cd codegen/apidocs/html; zip -rq  ../apidocs_html.zip *
cd -
go-bindata -pkg apidocs -prefix codegen/apidocs -o codegen/apidocs/apidocs_html_zip.go codegen/apidocs/apidocs_html.zip

# date files
cd codegen/date
go-bindata -pkg date -o bindate.go date_only.go  datetime.go datetime_only.go datetime_rfc2616.go  time_only.go
cd -

# build
godep go install -v ./...
