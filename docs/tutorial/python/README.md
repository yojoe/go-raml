# Python Tutorial

In this tutorial we will generate a simple Python server & client from [an RAML file](../api.raml) and then integrate it
with [itsyou.online](https://www.itsyou.online/) authorization server.


## Server

Generate server code by using this command

```
go-raml server -l python --ramlfile ../api.raml  --dir server
```

You can find all server files in `server` directory.

### Install required packages

```
$ cd server
$ virtualenv -p python3 env && source env/bin/activate && pip3 install -r requirements.txt
```

### Server side itsyou.online integration

We need to write/modify some code for this integration

**Install required packages**

```
$ sudo apt-get install libffi-dev
pip3 install pyjwt cryptography
```

**Modify generated Oauth2 middleware**

We need to modify `oauth2_itsyouonline.py` in order to integrate it with itsyou.online.

Add this code after last `import` statement

```
import jwt

iyo_pub_key = """
-----BEGIN PUBLIC KEY-----
MHYwEAYHKoZIzj0CAQYFK4EEACIDYgAES5X8XrfKdx9gYayFITc89wad4usrk0n2
7MjiGYvqalizeSWTHEpnd7oea9IQ8T5oJjMVH5cc0H5tFSKilFFeh//wngxIyny6
6+Vq5t5B0V0Ehy01+2ceEon2Y0XDkIKv
-----END PUBLIC KEY-----"""

token_prefix = "token "

def get_jwt_scopes(token):
    if token.startswith(token_prefix):
        token = token[len(token_prefix):]
        return jwt.decode(token, iyo_pub_key, verify=False)["scope"]
    else:
        raise Exception('invalid token')

```

And then changes these lines from
```
# provide code to check scopes of the access_token
scopes = []
```

to

```
# provide code to check scopes of the access_token
scopes = get_jwt_scopes(token)
```

You can find the modified file in [oauth2_itsyouonline.py](./server/oauth2_itsyouonline.py)

**execute the server**

```python3 app.py```


## Client

Generate client code by using this command

```
go-raml client --ramlfile ../api.raml --dir client -l python
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

**Install itsyou.online client library**

We need it to authenticate using JWT token from itsyou.online.

Because there is still no pypi package for itsyou.online client library,
we need to copy it manually

```
$ git clone --depth=1 https://github.com/itsyouonline/identityserver.git
$ cp -a identityserver/clients/python/itsyouonline .
$ rm -rf identityserver
```

**simple client main code**

A simple example of the client program can be found in [main.py](main.py).

The code is well commented to give you idea about what happens in the code.

**execute the client**

`python3 main.py [application_id] [application_secret]`

Change `[application_id]` and `[application_secret]` above with your own application ID and secret.