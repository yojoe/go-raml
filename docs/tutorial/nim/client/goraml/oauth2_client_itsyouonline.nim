import strutils

import client_goraml

let baseUri = "https://itsyou.online/v1/oauth/access_token"
proc getTokenByClientCrendentials*(c: Client, clientID: string, clientSecret: string, scopes: openArray[string], auds: openArray[string]):string =
  var q: seq[string]

  q = @[]
  q.add("grant_type=client_credentials")
  q.add("client_id=" & clientID)
  q.add("client_secret=" & clientSecret)
  q.add("response_type=id_token")

  if scopes.len > 0:
    q.add("scope=" & scopes.join(","))

  if auds.len > 0:
    q.add("aud=" & auds.join(","))

  return c.request(baseUri & "?" & q.join("&") , "POST").body