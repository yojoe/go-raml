#capnp

Table of Contents
=================

* [Install Capnp Tools](#install-capnp-tools)
* [Schema Generation](#schema-generation)
  * [Plain Schema](#plain-schema)
  * [Go-compatible schema](#go-compatible-schema)
* [Schema Compilation](#schema-compilation)

Other than source code, go-raml able to produce [canpnp schema](https://capnproto.org/language.html). 

go-raml can produce two kind of schemas:

- plain schema : without additional info for the compiler. It tested against C++(default), 
  [python](http://jparyani.github.io/pycapnp/), and [Nim](https://github.com/zielmicha/capnp.nim).
  This schema should work with other compilers which don't need metadatas.

- Go-compatible schema: capnp schema with additional infos for [go-capnproto2](https://github.com/zombiezen/go-capnproto2) compiler.

## install capnp tools

Unix (Linux, Mac, BSD)

```
curl -O https://capnproto.org/capnproto-c++-0.5.3.tar.gz
tar zxf capnproto-c++-0.5.3.tar.gz
cd capnproto-c++-0.5.3
./configure
make -j6 check
sudo make install
```
Other ways of installations can be found at https://capnproto.org/install.html

**Install Go compiler plugin**

You only need to do this if you want to use capnp with Go.

The generated capnp schema is compatible with Go compiler from zombiezen https://github.com/zombiezen/go-capnproto2
```
go get -u -t zombiezen.com/go/capnproto2/...
```

**Install python library**

You only need to do this if you want to use capnp with Python.

Installation procedures can be found at http://jparyani.github.io/pycapnp/install.html.
It is a library that use `.capnp` files directly than produce source code like other languages implementation.

**Install Nim compiler**

You only need to do this if you want to use capnp with Nim.

Installation procedures can be found at https://github.com/zielmicha/capnp.nim.

At the time of this writing, the compiler still has issue, please check https://github.com/zielmicha/capnp.nim/pull/2 for the fix.

Or simply use `nimble-fix` branch of https://github.com/iwanbk/capnp.nim


## Schema Generation

To generate capnp schema, go-raml supports an additional fields of RAML property:

- capnpType: if you want to specify [built-in types](https://capnproto.org/language.html#built-in-types) to use, you can use this field. It is **not mandatory** field.

Unless specified otherwise using `capnpType` field described above, go-raml use this type mapping

| RAML Type/Format             | Capnp Type|
| -----------------------------| ----------|
|string                        |  Text     |
|number                        | Float64   |
|integer                       | Int64     |
|boolean                       | Bool      |
| array                        | List      |
|file                          | Data      |
|object                        | Data      |
|datetime/date-only/time-only  | Text      |
|int8                          | Int8      |
|int16                         | Int16     |
| int64                        | Int64     |
|int                           | Int16     |
|float                         | Float64   |
|long                          | Int32     |
|double                        |Float64    |
 
 

### Plain schema

```
go-raml capnp --dir /result_dir --ramlfile api_file.raml -l plain
```

### Go-compatible schema

```
go-raml capnp --dir /tmp/wodaw/ --ramlfile codegen/capnp/fixtures/struct.raml -l go
```

## Schema Compilation

**Go**
```
capnp compile -I$GOPATH/src/zombiezen.com/go/capnproto2/std   -ogo *.capnp
```

**Nim**
```
capnp compile -onim *.capnp > file_result.nim
```

**Python**

We can use it directly, see this [quickstart guide](https://jparyani.github.io/pycapnp/quickstart.html).

