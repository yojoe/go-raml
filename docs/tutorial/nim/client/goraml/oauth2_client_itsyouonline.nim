import strutils, tables

import client_goraml

let baseUri = "https://itsyou.online/v1/oauth/access_token?response_type=id_token"
proc getAccessToken*(c: Client, clientID: string, clientSecret: string, scopes: openArray[string], auds: openArray[string]):string =
  var qp: Table[string, string] = {
    "grant_type": "client_credentials",
    "client_id": clientID,
    "client_secret": clientSecret
  }.toTable

  if scopes.len > 0:
    qp["scope"] = scopes.join(",")

  if auds.len > 0:
    qp["aud"] = auds.join(",")

  return c.request(baseUri, "POST", queryParams=qp).body