# Go Code Generator

## Server

Generated server code uses these libraries:
- [gorilla mux](https://github.com/gorilla/mux) as router
- [go-validator](https://github.com/go-validator/validator) for request body validation

## Client

Generated client library only use `http` package from stdlib.

## Type

RAML Object become Go struct, generated in `types` directory.

Because all types are generated into `types` package, its need to be exported
in order to be usable from other packages.
In other word, first letter of the type is going to be converted to uppercase.
It implies that it is invalid to have two types which only differ in the case(upper or lower)
of the first character.

Some rules about type's properties naming:

- capitalizing first character of the properties.
- json tag is the same as property name

#### Scalar Type Mapping
    Raml        |  Go   
    ----------- | -----------
    string      | string
    number      | float64
    integer     | int
    boolean     | bool
    date        | goraml.Date
    enum        | see below for explanation
    file        | string
    Array       | Array
    Union       | interface{}

### Enum

Enum is converted into:
- new type definition
- const variable for each enum value

Enum type and file name is started by `Enum`


## Input Validation

Input validation is done using [go-validator](https://github.com/go-validator/validator).

|    Validation              |    Go
|--------------------------- | ------
| minLength                  |   v
| maxLength                  |   v
| pattern                    |   v
| minimum                    |   v
| maximum                    |   v
| format                     |   x
| multipleOf                 |   v
| array field minItems       |   v
| array field maxItems       |   v
| array field uniqueItems    |   v
| array Type minItems        |   v
| array Type maxItems        |   v
| array Type uniqueItems     |   v


## Bodies
Request  and response body are mapped into structs
and following the same rules as types.

struct name = [Resource name][Method name][ReqBody|RespBody].


## Resources and Nested Resources

### Server
Resources in the server are mapped to:

- interface:
    - file name = [resource]_if.go
    - always be regenerated
    - interface name = [resource]Interface

- API implementation that implements the interface
    - main API file = handlers/[resource]/[resource]_api.go
    - per method API files also generated in the same directory as the main API file
    - only generated when the file is not present
    - struct name = [resource]API

- routes for all necessary routes:
    - func name = [Resource]InterfaceRoutes
    - generated in the same file as interface file

#### Catch-All route

`{path:*}` could be used to express Catch-All route.
For example /path/{path:*} is going to match:
  - /path/a    -> path = a
  - /path/a/b  -> path = a/b
  - /path/a/b/c -> path = a/b/c

### Client

Resourcess in the client are implemented as services.

Let's say we have two root resources:
- /users
- /network

Client library is going to have two services:
- Users
- Network

Each service will have it's own methods

## Methods

### Header

Code related to [Requests Headers](https://github.com/raml-org/raml-spec/blob/master/versions/raml-10/raml-10.md/#headers) are only generated in the Client lib. All functions have arguments to send any request header, the current client lib will not check the headers against the RAML specifications.


Response headers related code are only generated in the server in the form of commented code, example:
```
// uncomment below line to add header
// w.Header.Set("key","value")
```

### Query Strings and Query Parameters

All client library functions have arguments to send [query strings and query Parameters](https://github.com/raml-org/raml-spec/blob/master/versions/raml-10/raml-10.md/#query-strings-and-query-parameters).

The generated code in the server is in the form of commented code:

```
// name := req.FormValue("name")
```

## Responses

## Resource Types and Traits

[Resource Types and Traits](https://github.com/raml-org/raml-spec/blob/master/versions/raml-10/raml-10.md/#resource-types-and-traits) already parsed by the parser. So, the generator need to know nothing about it.

## Security Schemes

go-raml only supports [OAuth2.0](https://github.com/raml-org/raml-spec/blob/master/versions/raml-10/raml-10.md/#oauth-20).

- client : it currently able to get oauth2 token with client credentials.
- server : it currently only support JWT token from itsyou.online

## Annotations

## Modularization

### Includes

Includes should work properly

### Libraries

Library will be generated into a package.


### Overlays and Extensions

