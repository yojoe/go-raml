#Go-RAML Tutorial
1. export your GOPATH to your bashrc or .zhrc depending on what shell you are using

  ```
  $export GOPATH=/opt/go
  ```

  ```
  $export PATH=$PATH:$GOPATH/bin
  ```

2. Install go-raml

  ```
  $go get -u github.com/Jumpscale/go-raml
  ```

3. create dir for generated Go code

  ```
  $mkdir -p /opt/go/src/github.com/mycompany/server
  ```
  ##Code Generation
  we will use simple_example.raml exists in the raml/samples

4. generate server code GoLang

  ```
  $go-raml server -l go --dir /opt/go/src/github.com/mycompany/server --ramlfile tutorial/tutorial.raml --import-path github.com/mycompany/server
  ```
  
  
5. generate client code in python

  ```
  mkdir -p tutorial/client
  $go-raml client -l python --dir ./tutorial/client --ramlfile tutorial/tutorial.raml
  ```
  
Install required packages
  ```
  $cd /opt/go/src/github.com/mycompany/server
  $go get github.com/gorilla/mux
  $go get gopkg.in/validator.v2
  ```
  
  You can find more info about generated code and how to use it in https://github.com/Jumpscale/go-raml/blob/master/README.md#using-generated-code

  ## Playing with the generated code
  we will edit methods of the server/resources_api.go method with any dummy response
  for example for resourceIdGet 

  <code>fmt.Fprintf(w, "Actual implementation should return a resource")</code>

  to resourceIdGet method 


  and don't forget to <code>import fmt</code>


  ```
  $go build
  $./server
  ```

6. open a new terminal and cd to client dir under tutorial

  install ipython you didn't have it

  ```
  $pip install ipython
  ```
  Run ipython 

  ```
  $ipython
  ```

  ```python
      from client import  Client  
      c = Client()  
      c.url = "http://localhost:8080"  
      response = c.c.resources_get()  
      print (response.content)  
  ```

