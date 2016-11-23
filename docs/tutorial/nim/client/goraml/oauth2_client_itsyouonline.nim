import strutils

import client_goraml

let baseUri = "https://itsyou.online/v1/oauth/access_token?response_type=id_token"
proc getAccessToken*(c: Client, clientID: string, clientSecret: string, scopes: openArray[string], auds: openArray[string]):string =
  var q: seq[string]

  q = @[]
  q.add("grant_type=client_credentials")
  q.add("client_id=" & clientID)
  q.add("client_secret=" & clientSecret)

  if scopes.len > 0:
    q.add("scope=" & scopes.join(","))

  if auds.len > 0:
    q.add("aud=" & auds.join(","))

  var sep: string = "?"

  if baseUri.find("?") > 0:
    # there is already ? in the uri, we need to append "&" instead
    sep = "&"

  return c.request(baseUri & sep & q.join("&") , "POST").body
