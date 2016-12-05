# Nim Code Generator

## Server

Generated server code uses [jester](https://github.com/dom96/jester) as web framework.

## Client

Generated client library uses httpclient module from stdlib

## Type

RAML Object usually become Nim Object

#### Scalar Type Mapping
    Raml        |  Nim 
    ----------- | -----------
    string      | string
    number      | float64
    integer     | int
    boolean     | bool
    date        | Time
    enum        | enum
    file        | string
    Array       | sequence
    Union       | -


## Input Validation

TBD


## Bodies
Request  and response body are mapped into structs
and following the same rules as types.

struct name = [Resource name][Method name][ReqBody|RespBody].


## Resources and Nested Resources

### Server

Resources in the server are mapped to:

- routes in main file:
    
- API implementation that implements the resource. One file for each root resource

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

TBD


### Query Strings and Query Parameters

All client library functions have arguments to send [query strings and query Parameters](https://github.com/raml-org/raml-spec/blob/master/versions/raml-10/raml-10.md/#query-strings-and-query-parameters).


## Responses

## Resource Types and Traits

[Resource Types and Traits](https://github.com/raml-org/raml-spec/blob/master/versions/raml-10/raml-10.md/#resource-types-and-traits) already parsed by the parser. So, the generator need to know nothing about it.

## Security Schemes

go-raml only supports [OAuth2.0](https://github.com/raml-org/raml-spec/blob/master/versions/raml-10/raml-10.md/#oauth-20).

- client : it currently able to get oauth2 token with client credentials.
- server : it currently only support JWT token.

## Annotations

## Modularization

### Includes

Includes should work properly

### Libraries

-

### Overlays and Extensions

