import marshal
import client

import City


type
  Users_service* = object
    client*: Client
    name*: string

proc UsersSrv*(c : Client) : Users_service  =
  return Users_service(client:c, name:c.baseURI)


proc GetUsers*(srv: Users_service) : usersGetRespBody =
  let resp = srv.client.request("/users", "GET")
  return to[usersGetRespBody](resp.body)

proc usersPost*(srv: Users_service, reqBody: City) : City =
  let resp = srv.client.request("/users", "POST", $$reqBody)
  return to[City](resp.body)

proc OptionsUsers*(srv: Users_service) : string =
  let resp = srv.client.request("/users", "OPTIONS")
  return to[string](resp.body)

proc GetUserByID*(srv: Users_service, userId: string) : City =
  let resp = srv.client.request("/users/"&userId, "GET")
  return to[City](resp.body)

proc usersByUserIdDelete*(srv: Users_service, userId: string) : string =
  let resp = srv.client.request("/users/"&userId, "DELETE")
  return to[string](resp.body)

proc usersByUserIdAddressPost*(srv: Users_service, reqBody: usersuserIdaddressPostReqBody, userId: string) : usersuserIdaddressPostRespBody =
  let resp = srv.client.request("/users/"&userId&"/address", "POST", $$reqBody)
  return to[usersuserIdaddressPostRespBody](resp.body)

proc usersByUserIdAddressFolderAddressIdTestAddressId2Get*(srv: Users_service, addressId: string, addressId2: string, userId: string) : seq[address] =
  let resp = srv.client.request("/users/"&userId&"/address/folder"&addressId&"test"&addressId2, "GET")
  return to[seq[address]](resp.body)

