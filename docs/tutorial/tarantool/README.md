# Tarantool Tutorial

In this tutorial we will generate a simple tarantool server from [an RAML file](../api.raml)

## Server

Generate tarantool server code by using this command

```
go-raml server -l tarantool --ramlfile ../api.raml  --dir server
```

You can find all tarantool files in `server` directory.

**execute the server**

```tarantool main.lua```
