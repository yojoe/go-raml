# Nim Tutorial

In this tutorial we will generate a simple NIM server & client from [an RAML file](../api.raml).

## Server

Generate server code by using this command

```
go-raml server -l nim --ramlfile ../api.raml  --dir server
```
### install libjwt

The server needs [libjwt](https://github.com/benmcollins/libjwt) for JWT decoding.

Install it using provided script (for Debian/Ubuntu)

```
./install_libjwt.sh
```

### Compile and Run it

```
./run.sh
```

You can see the server by pointing your browser to `http://localhost:5000/`

API Docs can be seen by clicking `API Docs` link on that page.

## Client

Generate goraml client code by using this command

```
go-raml client --ramlfile ../api.raml --dir client/goraml -l nim
```
Then you can find goraml client code in `client/goraml` directory.


Generate itsyouonline client lib by using this command

```
go-raml client -l nim --ramlfile ../itsyouonline-raml/itsyouonline.raml --dir client/iyo 
```
It will generate itsyouonline client library in `client/iyo' directory


### Simple usage

**client**

There is `main.nim` in `client` directory. It serves as an example on 
how to use generated client lib.

You need to modify it by supply your itsyouonline client-id and client-secret
when calling `createJWTToken` proc

**server**

`users_api.nim` file in `server` directory has been modified to
serve the client.

Added lines has `added line` comment in the end of line.

**run the client example**

```
./run.sh
```
