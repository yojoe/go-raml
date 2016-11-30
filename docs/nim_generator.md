# Nim Code Generator

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

Resources in the server are mapped to:

- routes in main file:
    
- API implementation that implements the interface

## Methods

### Header

TBD


### Query Strings and Query Parameters

TBD

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

Libraries should work properly except the apidocs web page (REST UI).

### Overlays and Extensions

