# Python Code Generator

## Requirements

- [autopep8](https://pypi.python.org/pypi/autopep8) for automatic code formatting

## Server

Generated server code use these libraries:

- Flask or Sanic as web framework
- [python jsonschema](https://pypi.python.org/pypi/jsonschema) for request body validation
- [python-jose](https://github.com/mpdavis/python-jose) for JWT decoding

## Client

Generated client library use [requests](http://docs.python-requests.org/en/master/) or 
[aiohttp](http://aiohttp.readthedocs.io/en/stable/) as http library.


## Type

RAML Object is mapped to python class


#### Scalar Type Mapping

-

### Enum

RAML Enum become python enum as described in https://docs.python.org/3/library/enum.html

## Input Validation

Except for builtin types, generated server use python-jsonschema for request body validation.


 |   Validation               | Status
 | -------------------------- | ------
 | minLength                  |   v
 | maxLength                  |   v
 | pattern                    |   v
 | minimum                    |   v
 | maximum                    |   v
 | format                     |   x
 | multipleOf                 |   v
 | array field minItems       |   v
 | array field maxItems       |   v
 | array field uniqueItems    |   x
 | array Type minItems        |   x
 | array Type maxItems        |   x
 | array Type uniqueItems     |   x


## Bodies
Request  and response body are mapped into structs
and following the same rules as types.

struct name = [Resource name][Method name][ReqBody|RespBody].


## Resources and Nested Resources

### Server

Resources in the server are mapped to:

- a flask blueprint module

The generated server contains a handlers directory, the python handlers files are the only files that are not overwritten by a regeneration. Each api method has a generated handler file, this is where the user should add the implementation for each method.

### Client

Resourcess in the client are implemented as services.

Let's say we have two root resources:
- /users
- /network

Client library is going to have two services:
- users
- network

Each service will have it's own methods


## Methods

### Header

Code related to [Requests Headers](https://github.com/raml-org/raml-spec/blob/master/versions/raml-10/raml-10.md/#headers) are only generated in the Client lib. All functions have arguments to send any request header, the current client lib will not check the headers against the RAML specifications.



### Query Strings and Query Parameters

All client library functions have arguments to send [query strings and query Parameters](https://github.com/raml-org/raml-spec/blob/master/versions/raml-10/raml-10.md/#query-strings-and-query-parameters), the current client lib will not check it against the RAML specifications.


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

Libraries should work properly except the apidocs web page (REST UI).


### Overlays and Extensions

