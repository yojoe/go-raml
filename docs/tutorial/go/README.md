# Go Tutorial

In this tutorial we will generate a simple Go client & server from [an RAML file](../api.raml) and then integrate it
with [itsyou.online](https://www.itsyou.online/) authorization server.

This tutorial use [jumpscale/ubuntu1604_golang](https://hub.docker.com/r/jumpscale/ubuntu1604_golang/) docker image.
But you can use any system with Go installed and configured.

## Create docker instances

You can skip this step if you want to use your own enviroment.

```
sudo docker pull jumpscale/ubuntu1604_golang
sudo docker run --rm -t -i  --name=goraml jumpscale/ubuntu1604_golang
```

## Server

Generate server code by using this command
```
$ go-raml server --ramlfile ../api.raml --import-path examples.com/goramldir --dir $GOPATH/src/examples.com/goramldir
```

You can find all server files in `$GOPATH/src/examples.com/goramldir` directory.


### Server side itsyou.online integration

You only need to replace the value of `oauth2ServerPublicKey` variable in `oauth2_itsyouonline_middleware.go`
to the content of [itsyouonline.pub](../itsyouonline.pub).


**Build & Run the server**
```
go build
./goramldir
```

## Client

generate client code by using this command

`go-raml client --ramlfile ../api.raml --dir client/goramldir --package goramldir`

Then you can find client code in `client` directory.


**simple client main program**

A simple example of the client program can be found in [main.go](client/main.go).

Steps to use generated client lib:

- create `goramldir` client object
- create itsyou.online JWT token
- set JWT token as authorization header

after above steps, client are ready to make API call to `goramldir` server.

The code is relatively simple and have enough comment, so it should be easy to understand. 

**execute client program**

```
go build
./client --app_id=YOUR_APP_ID --app_secret=YOUR_API_KEY
```

## Auto Generated API server homepage & API Docs

`go-raml` generated code also provide you with a simple homepage and auto generated API docs.

To access it, you can visit http://localhost:5000 from your browser

![API server homepage](images/api_home.png)

If we click `API Docs`, browser will go to auto generated API docs page

![API Docs](images/apidocs_home.png)

We can click any HTTP verbs in the API Docs page to see documentation for that endpoint & verb.
This is documentation for `POST /users` page

![POST /users](images/apidocs_post_details.png)
