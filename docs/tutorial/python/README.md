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

Generate `requests` based client code by using this command

```
go-raml client --ramlfile ../api.raml --dir requests_client/goramldir -l python
```
Then you can find client code in `requests_client` directory.
You will need to install `requests` library to use the client.


You can also generate `aiohttp` asyncio based client code by using this command

```
go-raml client --ramlfile ../api.raml --dir aiohttp_client/goramldir -l python -kind aiohttp
```

Then you can find client code in `aiohttp_client` directory.
You will need to install `aiohttp` and `aiodns` library to use the client.


**simple client main code**

A simple example of the client program can be found in main.py.

The code is well commented to give you idea about what happens in the code.

**execute the client**

`python3 main.py [application_id] [application_secret]`

Change `[application_id]` and `[application_secret]` above with your own application ID and secret.
