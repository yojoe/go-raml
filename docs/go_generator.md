# Go Code Generator

## Type

RAML Object usually become Go struct.

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


    Validation              |    Go | Python
--------------------------- | ------| -----------
 minLength                  |   v   |   v
 maxLength                  |   v   |   v
 pattern                    |   v   |   v
 minimum                    |   v   |   v
 maximum                    |   v   |   v
 format                     |   x   |   x
 multipleOf                 |   v   |   v
 array field minItems       |   v   |   v
 array field maxItems       |   v   |   v
 array field uniqueItems    |   v   |   x
 array Type minItems        |   v   |   x
 array Type maxItems        |   v   |   x
 array Type uniqueItems     |   v   |   x


## Bodies
Request  and response body are mapped into structs
and following the same rules as types.

struct name = [Resource name][Method name][ReqBody|RespBody].


## Resources and Nested Resources

Resources in the server are mapped to:

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


## Methods

### Header

Code related to [Requests Headers](https://github.com/raml-org/raml-spec/blob/master/versions/raml-10/raml-10.md/#headers) are only generated in the Client lib. All functions have arguments to send any request header, the current client lib will not check the headers against the RAML specifications.


Response headers related code are only generated in the server in the form of commented code, example:
```
// uncomment below line to add header
// w.Header.Set("key","value")
```

### Query Strings and Query Parameters

All client library functions have arguments to send [query strings and query Parameters](https://github.com/raml-org/raml-spec/blob/master/versions/raml-10/raml-10.md/#query-strings-and-query-parameters), the current client lib will not check it against the RAML specifications.

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

Libraries should work properly except the apidocs web page (REST UI).


### Overlays and Extensions

