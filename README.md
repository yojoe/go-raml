# go-raml
[![Build Status](https://travis-ci.org/Jumpscale/go-raml.svg?branch=master)](https://travis-ci.org/Jumpscale/go-raml)


Table of Contents
=================

* [What is go-raml](#what-is-go-raml)
* [Supported RAML Versions](#raml-versions)
* [Install & Build in Development](#install)
* [Code Generation](#code-generation)
* [Usage](#usage)
* [Generating Server](#generating-server)
  * [Simple Homepage & API Docs](#simple-home-page-and-api-docs)
  * [Go Server](#go-server)
  * [Flask / Python Server](#flaskpython-server)
* [Generating Client](#generating-client)
* [Specification File](#specification-file)
* [Contribute](#contribute)
* [Roadmap](#roadmap)
* [RAML to Code Translation](#raml-to-code-translation)


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

Currently there are still some [limitations](docs/limitations.md) on the RAML 1.0 features that are supported.

## Install

make sure you have at least go 1.6 installed !

`go get -u github.com/Jumpscale/go-raml`


### Build in development

Install `godep` as package manager, we need it for vendoring tool

`$go get -u github.com/tools/godep`

Install go-bindata, we need it to compile all resource files to .go file

`go get -u github.com/jteeuwen/go-bindata/...`


To rebuild APIDocs files to .go file
```
cd $GOPATH/src/github.com/Jumpscale/go-raml
sh build_apidocs.sh
```

Build go-raml and all resource files
```
cd $GOPATH/src/github.com/Jumpscale/go-raml
sh build.sh
```

## Code generation

Internally, go templates are used to generate the code, this provides a flexible way to alter the generated code and to add different languages for the client.

## Usage

To use it on the commandline, just execute `go-raml` without any arguments, it will output the help on the stdout.

To invoke the codegeneration using `go generate`, specify the generation in 1 of your go source files:
`//go:generate go-raml ...`

go-raml needs to be on the path for this to work off course.


## Generating Server

`go-raml` able to generate Go & Python server.

Generated server will listen on port 5000


### Go Server

To generate the Go code for implementing the server in first design approach, execute

`go-raml server -l go --dir /gopath/src/github.com/mycompany/myapi --ramlfile api.raml --import-path github.com/mycompany/myapi`

`--dir` specify generated code directory. Except you really know what you do, use $GOPATH as root directory. In above example $GOPATH is /gopath.


`--ramlfile` is the path RAML file

`--import-path` is the root import path of the code. Generated code contains sub packages and need it to properly import code from another packages



### Flask/Python Server

To generate the Flask/Python code for implementing the server, execute

`go-raml server -l python --dir ./result_directory --ramlfile api.raml`

### Code Generator Options

```
   --language, -l "go"  Language to construct a server for
   --dir "."        target directory
   --ramlfile "."   Source raml file
   --package "main" package name
   --no-main        Do not generate a main.go file
   --no-apidocs     Do not generate API Docs in /apidocs/?raml=api.raml endpoint
   --import-path    "examples.com/ramlcode"	import path of the generated code
```

## Generating Client

`go-raml client --language go  --dir ./result_directory --ramlfile api.raml`

A go 1.5.x compatible client is generated in result_directory directory.

`go-raml client --language python --dir ./result_directory --ramlfile api.raml`

A python 3.5 compatible client is generated in result_directory directory.


## Using Generated Code

### Simple home page and API Docs

Both servers have simple HTML home page that can be accessed in http://localhost:5000.
it provide simple description of the API server and link to auto generated API Documentation
powered by [api-console](https://github.com/mulesoft/api-console)

### Using Go Server

Generated go server can be used as your API server skeleton. It gives you routing code, implementation skeleton, simple request validation,  and all struct based on raml types and request/response body. You can then extend the generated code to fit your need.

Generated code details:

- an optional main file which contain routing code and other inialization code
- Interfaces files, contain routing code which mapped one to one with RAML resources. It always regenerated
- Interface implementation. It is code skeleton of the resource implemetation, only generated when the file is not present. So you can easily modify it
- Struct for RAML types and request/response body, only generated when the file is not present.
- Validation code
- Helper files

The generated server uses [Gorilla Mux](http://www.gorillatoolkit.org/pkg/mux) as HTTP request multiplexer.

Build the code

`go build ./...`

Execute it

`./binary_name`

The server will then run in port 5000, you can go to http://localhost:5000 to see default html page as described above

### Using generated python server

The generated code is utilizing [Flask Blueprint](http://flask.pocoo.org/docs/0.11/blueprints/) to give you modular flask code.

Generated code details:

- main file named app.py which initialize the app
- one bluperint/module for each root RAML resource
- helper file
- requirements.txt which contains list of packages needed to run the app.

Install needed packages
```
pip3 install -r requirements.txt
```

You might want to install it inside virtualenv

Run the code
```
python3 app.py
```

### Using Go client library

Generated go client library can be used directly as modular package of your project

### Using Python Client library

Generated python client library only need python-requests as dependency

## Specification file

Besides generation of a new RAML specification file, updating an existing raml file is also supported. This way the raml filestructure that can be included in the main raml file is honored.

`go-raml spec ...`

## Contribute

When you want to contribute to the development, follow the [contribution guidelines](contributing.md).

## Roadmap
**v0.1**

* Generation of the server using [gorilla mux](http://www.gorillatoolkit.org/pkg/mux)
* Generation of a go client
* Generation of a python 3.5 client
* Generation of a python flask server

**v0.2**

* OAuth 2.0 support
* Input validation according to the RAML type definitions

**v0.3**

* Possibility to generate a default server implementation to MongoDB

**v0.4**

* Generation of a new RAML specification file

**v0.5**

* Update of a RAML specification file


## RAML to Code Translation

Below are incomplete description about how we translate .raml file into code.

### Type

[types](http://docs.raml.org/specs/1.0/#raml-10-spec-types) is mapped to struct.

Some rules about properties naming:

- capitalizing first character of the properties.
- json tag is the same as property name

File name is the same as types name with lowercase letter.
struct name = types name.

#### Type Mapping
    Raml        |  Go   
    ----------- | -----------
    string      | string
    number      | float
    integer     | int
    boolean     | bool
    date        | Date
    enum        | []string
    file        | string
    sometype[]  | []sometype
    sometype[][]| [][]sometype
    Union       | interface{}

### Bodies
[Request Body](http://docs.raml.org/specs/1.0/#raml-10-spec-bodies) and response body are mapped into structs
and following the same rules as types.

struct name = [Resource name][Method name][ReqBody|RespBody].

RequestBody generated from body node below method.

ResponseBody generated from body node below responses.

### Resource
[Resource](http://docs.raml.org/specs/1.0/#raml-10-spec-resources-and-nested-resources) in the server is mapped to:
- interface:

    - file name = [resource]_if.go
    - always be regenerated
    - interface name = [resource]Interface

- API implementation that implements the interface
    - file name = [resource]_api.go
    - only generated when the file is not present
    - struct name = [resource]API

- routes for all necessary routes:
    - func name = [Resource]InterfaceRoutes


### Header

Code related to requests [headers](http://docs.raml.org/specs/1.0/#raml-10-spec-headers) are only generated in the Client lib. All functions have arguments to send any request header, the current client lib will not check the headers against the RAML specifications.


Response headers related code is only generated in the server in the form of commented code, example:
```
// uncomment below line to add header
// w.Header.Set("key","value")
```

### Query Strings and Query Parameters

All client library functions have arguments to send [Query Strings and Query Parameters](http://docs.raml.org/specs/1.0/#raml-10-spec-query-strings-and-query-parameters), the current client lib will not check it against the RAML specifications.

The generated code in the server is in the form of commented code:

```
// name := req.FormValue("name")
```

### Input Validation


    Validation      |    Go     | Python
------------------------------- | ------------- | -----------
 minLength                      |   v   |   v
 maxLength          |   v   |   v
 pattern            |   v   |   v
 minimum            |   v   |   v
 maximum            |   v   |   v
 format             |       x   |   x
 multipleOf         |   v   |   v
 array field minItems       |   v   |   v
 array field maxItems       |   v   |   v
 array field uniqueItems    |   v   |   x
 array Type minItems        |   v   |   x
 array Type maxItems        |   v   |   x
 array Type uniqueItems     |   v   |   x
