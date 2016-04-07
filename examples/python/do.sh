mkdir -p server
go-raml server -l python --dir ./server --ramlfile api.raml
go-raml client -l python --dir ./client --ramlfile api.raml
