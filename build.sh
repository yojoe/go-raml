#!/bin/bash
./build_template.sh

# apidocs.zip
cd codegen/apidocs/html; zip -rq  ../apidocs_html.zip *
cd -
go-bindata -pkg apidocs -prefix codegen/apidocs -o codegen/apidocs/apidocs_html_zip.go codegen/apidocs/apidocs_html.zip

# build
godep go install -v ./...
