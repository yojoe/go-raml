# Tarantool Code Generator

## Server

The server is run using [tarantool](https://tarantool.org/en/download/download.html)

Generated server code uses these libraries:
- [tarantool-http](https://github.com/tarantool/http/) as router

The generator uses the following libraries:
- [capnpproto](https://github.com/capnproto/capnproto) to compile capnp
- [lua-capnpproto](https://github.com/cloudflare/lua-capnproto) to generate lua from compiled capnp


to generate a server run the following command:
```bash
go-raml server -l tarantool --ramlfile <ramlfile> --dir <destination>
```

it will generate server code in <destination> for raml located at <ramlfile>.

the generator will generate capnp files for all the raml types then compile it into a lua file

### Structure
- main.lua: the main server file, it contains all the initialization and start of the server as well as all registering all the routes.
- handlers: a directory containing all the routes handlers. For each end point there is one handler file which contains a handler function for each request method.
- schemas: a directory containing all the generated capnp files and schema.lua which contains the compiled lua.

Running `tarantool main.lua` will start the tarantool server.