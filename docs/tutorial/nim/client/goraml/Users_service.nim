import marshal
import client_goraml

import User


type
  Users_service* = object
    client*: Client
    name*: string

proc UsersSrv*(c : Client) : Users_service  =
  return Users_service(client:c, name:c.baseURI)


proc usersGet*(srv: Users_service) : seq[User] =
  let resp = srv.client.request("/users", "GET")
  return to[seq[User]](resp.body)

proc usersPost*(srv: Users_service, reqBody: User) : User =
  let resp = srv.client.request("/users", "POST", $$reqBody)
  return to[User](resp.body)

proc usersByUsernameGet*(srv: Users_service, username: string) : User =
  let resp = srv.client.request("/users/"&username, "GET")
  return to[User](resp.body)

