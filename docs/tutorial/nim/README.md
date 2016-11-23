# Nim Tutorial

In this tutorial we will generate a simple NIM server & client from [an RAML file](../api.raml).

## Server

Generate server code by using this command

```
go-raml server -l nim --ramlfile ../api.raml  --dir server
```
Copy itsyou.online public key

```
cp ../itsyouonline.pub server/oauth2_server_key.pub
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


### Simple usage

**client**

There is `main.nim` in `client` directory. It serves as an example on 
how to use generated client lib.

You need to modify these:
- itsyouonline clientId and clientSecret variables

The example call an endpoint which doesn't need any scope.
So, as long as you have valid itsyouonline token, you can call it.

There are commented lines of code, which call endpoint that need scopes. 
You need to uncomment thoselines and modify these:
- "user:memberof:goraml" to your organization/scope
- "external1" to your audience

**server**

`users_api.nim` file in `server` directory has been modified to
serve the client.

Added lines has `added line` comment in the end of line.

**run the client example**

```
./run.sh
```
