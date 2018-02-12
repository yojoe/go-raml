package gentest

//go:generate go-raml server -l python --ramlfile ../../fixtures/gentest/goramldir.raml  --dir goramldir/server --no-apidocs
//go:generate go-raml server -l python --ramlfile ../../fixtures/gentest/goramldir.raml  --dir goramldir/sanic --kind=sanic --no-apidocs
