import marshal, tables
import client_goraml

import User


type
  Users_service* = object
    client*: Client
    name*: string

proc UsersSrv*(c : Client) : Users_service  =
  return Users_service(client:c, name:c.baseURI)


proc usersGet*(srv: Users_service, queryParams: Table[string, string] = initTable[string, string]()) : seq[User] =
  let resp = srv.client.request("/users", "GET", queryParams=queryParams)
  return to[seq[User]](resp.body)

proc usersPost*(srv: Users_service, reqBody: User, queryParams: Table[string, string] = initTable[string, string]()) : User =
  let resp = srv.client.request("/users", "POST", $$reqBody, queryParams=queryParams)
  return to[User](resp.body)

proc usersByUsernameGet*(srv: Users_service, username: string, queryParams: Table[string, string] = initTable[string, string]()) : User =
  let resp = srv.client.request("/users/"&username, "GET", queryParams=queryParams)
  return to[User](resp.body)

