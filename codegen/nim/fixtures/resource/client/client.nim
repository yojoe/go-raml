import httpclient

type
  Client* = object
    baseURI*: string
    hc: HttpClient

const defaultBaseURI = "http://localhost:8080"

proc newClient*(baseURI = defaultBaseURI): Client =
  # creates new client
  var c = Client(baseURI: baseURI, hc: newHttpClient())
  return c

proc request*(c: Client, endpoint: string, httpMethod = "GET", body = ""): httpclient.Response =
  return c.hc.request(c.baseURI & endpoint, httpMethod, body)