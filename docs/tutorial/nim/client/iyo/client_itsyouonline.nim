import httpclient, strutils

type
  Client* = object
    baseURI*: string
    hc: HttpClient

const defaultBaseURI = "https://itsyou.online/api"

proc newClient*(baseURI = defaultBaseURI): Client =
  # creates new client
  var c = Client(baseURI: baseURI, hc: newHttpClient())
  c.hc.headers = newHttpHeaders({ "Content-Type": "application/json" })
  return c

proc setAuthHeader*(c: Client, value: string) =
  c.hc.headers = newHttpHeaders({ "Content-Type": "application/json" })
  c.hc.headers.add("Authorization", value)

proc request*(c: Client, endpoint: string, httpMethod = "GET", body = ""): httpclient.Response =
  var uri: string = endpoint
  if not uri.startsWith("http"):
    uri = c.baseURI & uri
  return c.hc.request(uri, httpMethod, body)