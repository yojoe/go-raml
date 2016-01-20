# go-raml

## What is go-raml

When creating and maintaining api's, there are two approaches:
* Design first

    You design the api, all methods, descriptions and types from the api point of view. Afterwards, you generate all the boilerplate code and documentation to bootstrap the code.

* Code first

    When modifying the implementation, the interface definitions need to be kept in sync. Besides the good practice of keeping your specification up to date, it is used by other tools to generate clients, documentation, ...

This tool supports both (or at least, this is on the roadmap).
As a specification format, it uses [RAML 1.0](http://raml.org) .

## RAML versions
Only RAML version 1.0 RC is supported.


## Usage

go-raml is a commandline tool. To include the codegeneration in the go build, specify the generation in 1 of your go source files:
`//go:generate go-raml ...`

go-raml needs to be on the path for this to work off course.

To use it on the commandline yourself, just execute `go-raml` without any arguments, it will output the help on the stdout:


## Code generation

Internally, go templates are used to generate the code, this provides a flexible way to alter the generated code and to add different languages for the client.

## Server

To generate the go code for implementing the server in a design first approach, execute

`go-raml server ...`





## Client

`go-raml client --language go ...`

A go 1.5.x compatible client is generated.

`go-raml client --language python ...`

A python 3.5 compatible client is generated.

## Specification file

Besides generation of a new RAML specification file, updating an existing raml file is also supported. This way the raml filestructure that can be included in the main raml file is honored.

## roadmap
**v0.1**

* Generation of the server
* Generation of a go client
* Generation of a python 3.5 client

**v0.2**

* OAuth 2.0 support
* Generation of a new RAML specification file

**v0.3**

* Update of a RAML specification file
