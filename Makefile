PACKAGES = $(shell go list ./... | grep -v '/vendor/')

all: install

install:
	go generate
	go install -v

test:
	go test $(PACKAGES)
