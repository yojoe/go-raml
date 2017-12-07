package main

//go:generate go-bindata -pkg templates -nometadata  -prefix codegen -o codegen/templates/bindata.go codegen/templates/ codegen/templates/python codegen/templates/golang codegen/templates/nim codegen/templates/capnp codegen/templates/tarantool
//go:generate go-bindata -nometadata -pkg date -prefix codegen/date -o codegen/date/bindate.go codegen/date/date_only.go  codegen/date/datetime.go codegen/date/datetime_only.go codegen/date/datetime_rfc2616.go  codegen/date/time_only.go
