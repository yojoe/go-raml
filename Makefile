PACKAGES = $(shell go list ./... | grep -v '/vendor/')

all: install

install:
	go generate
	go install -v

apidocs:
	cd codegen/apidocs/html
	rm -f ../apidocs_html.zip
	zip -rq  ../apidocs_html.zip *
	cd -
	go-bindata -nometadata -pkg apidocs -prefix codegen/apidocs -o codegen/apidocs/apidocs_html_zip.go codegen/apidocs/apidocs_html.zip

test:
	go test $(PACKAGES)
