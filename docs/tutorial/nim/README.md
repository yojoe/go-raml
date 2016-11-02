# Nim Tutorial

In this tutorial we will generate a simple NIM server & client from [an RAML file](../api.raml).

## Server

Generate server code by using this command

```
go-raml server -l nim --ramlfile ../api.raml  --dir server
```

### Compile and Run it

```
nim c -r main.nim users_api.nim User.nim 
```

You can see the server by pointing your browser to `http://localhost:5000/`

API Docs can be seen by clicking `API Docs` link on that page.

## Client

Generate client code by using this command

```
go-raml client --ramlfile ../api.raml --dir client -l nim
```
Then you can find client code in client directory.

### Simple usage

**client**

There is `main.nim` in `client` directory. It serves as an example on 
how to use generated client lib.

**server**

`users_api.nim` file in `server` directory has been modified to
serve the client.

Added lines has `added line` comment in the end of line.

**run the client example**

```
nim c -r main.nim client.nim Users_service.nim User.nim 
```
