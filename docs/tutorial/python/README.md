# Python Tutorial

In this tutorial we will generate a simple Python server & client from [an RAML file](../api.raml) and then integrate it
with [itsyou.online](https://www.itsyou.online/) authorization server.


## Server

Generate Flask server code by using this command

```
go-raml server -l python --ramlfile ../api.raml  --dir flask
```

You can find all Flask files in `flask` directory.

Generate a Gevent Flask server code by using this command

```
go-raml server -l python --ramlfile ../api.raml  --dir gevent --kind gevent-flask
```

You can find all Flask files in `gevent` directory.

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


Generate `gevent requests` based client code by using this command

```
go-raml client --ramlfile ../api.raml --dir gevent_client/goramldir -l python --kind gevent-requests
```
Then you can find client code in `gevent_client` directory.
You will need to install `requests` and `gevent` library to use the client.


You can also generate `aiohttp` asyncio based client code by using this command

```
go-raml client --ramlfile ../api.raml --dir aiohttp_client/goramldir -l python --kind aiohttp
```

Then you can find client code in `aiohttp_client` directory.
You will need to install `aiohttp` and `aiodns` library to use the client.

### Unmarshall response

Client generator has an option to generate client code that unmarshall the response
body into the generated client class.
It is currently disabled by default and could be enabled by setting this option `--python-unmarshall-response` in the `go-raml client` command line option.

```
go-raml client --ramlfile ../api.raml --dir requests_client/goramldir -l python --python-unmarshall-response
```


In case of success unmarshal, the generated client returns two values:
- data : python class constructed by unmarshalling the response body
- response: python-requests's response object

otherwise it returns exception.

Usage example:

```python
try:
    data, resp = client.network.getNetwork("my_net_id")
    print(resp)
except unmarshall_error.UnmarshallError as ue:
    print("response:", ue.response.text)
    print("msg: ", ue.message)
```


**simple client main code**

A simple example of the client program can be found in main.py.

The code is well commented to give you idea about what happens in the code.

**execute the client**

`python3 main.py [application_id] [application_secret]`

Change `[application_id]` and `[application_secret]` above with your own application ID and secret.
