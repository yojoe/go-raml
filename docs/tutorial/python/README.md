# Python Tutorial

In this tutorial we will generate a simple Python server & client from [an RAML file](../api.raml) and then integrate it
with [itsyou.online](https://www.itsyou.online/) authorization server.


## Server

Generate Flask server code by using this command

```
go-raml server -l python --ramlfile ../api.raml  --dir flask
```

You can find all Flask files in `flask` directory.

Generate Sanic code by using this command

```
go-raml server -l python --ramlfile ../api.raml  --dir sanic --kind sanic
```

You can find all Sanic files in `sanic` directory.


### Install required packages

```
$ cd server
$ virtualenv -p python3 env && source env/bin/activate && pip3 install -r requirements.txt
```

### Server side itsyou.online integration

We need to write/modify some code for this integration


**Modify generated Oauth2 middleware**

Fill `oauth2_server_pub_key` variable with content [itsyouonline.pub](../itsyouonline.pub)

**execute the server**

```python3 app.py```


## Client

Generate client code by using this command

```
go-raml client --ramlfile ../api.raml --dir client/goramldir -l python
```
Then you can find client code in client directory.

**Install python requests library**

We need it to make HTTP request to server

```
$ cd client
$ virtualenv -p python3 env
$ source env/bin/activate
$ pip3 install requests
```


**simple client main code**

A simple example of the client program can be found in [main.py](main.py).

The code is well commented to give you idea about what happens in the code.

**execute the client**

`python3 main.py [application_id] [application_secret]`

Change `[application_id]` and `[application_secret]` above with your own application ID and secret.
