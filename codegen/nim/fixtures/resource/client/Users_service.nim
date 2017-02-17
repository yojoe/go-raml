import marshal, tables
import client_struct

import City
import address
import usersGetRespBody
import usersuserIdaddressPostReqBody
import usersuserIdaddressPostRespBody


type
  Users_service* = object
    client*: Client
    name*: string

proc UsersSrv*(c : Client) : Users_service  =
  return Users_service(client:c, name:c.baseURI)


proc GetUsers*(srv: Users_service, queryParams: Table[string, string] = initTable[string, string]()) : usersGetRespBody =
  let resp = srv.client.request("/users", "GET", queryParams=queryParams)
  return to[usersGetRespBody](resp.body)

proc usersPost*(srv: Users_service, reqBody: City, queryParams: Table[string, string] = initTable[string, string]()) : City =
  let resp = srv.client.request("/users", "POST", $$reqBody, queryParams=queryParams)
  return to[City](resp.body)

proc OptionsUsers*(srv: Users_service, queryParams: Table[string, string] = initTable[string, string]()) : string =
  let resp = srv.client.request("/users", "OPTIONS", queryParams=queryParams)
  return to[string](resp.body)

proc GetUserByID*(srv: Users_service, userId: string, queryParams: Table[string, string] = initTable[string, string]()) : City =
  let resp = srv.client.request("/users/"&userId, "GET", queryParams=queryParams)
  return to[City](resp.body)

proc usersByUserIdDelete*(srv: Users_service, userId: string, queryParams: Table[string, string] = initTable[string, string]()) : string =
  let resp = srv.client.request("/users/"&userId, "DELETE", queryParams=queryParams)
  return to[string](resp.body)

proc usersByUserIdAddressPost*(srv: Users_service, reqBody: usersuserIdaddressPostReqBody, userId: string, queryParams: Table[string, string] = initTable[string, string]()) : usersuserIdaddressPostRespBody =
  let resp = srv.client.request("/users/"&userId&"/address", "POST", $$reqBody, queryParams=queryParams)
  return to[usersuserIdaddressPostRespBody](resp.body)

proc usersByUserIdAddressFolderAddressIdTestAddressId2Get*(srv: Users_service, addressId: string, addressId2: string, userId: string, queryParams: Table[string, string] = initTable[string, string]()) : seq[address] =
  let resp = srv.client.request("/users/"&userId&"/address/folder"&addressId&"test"&addressId2, "GET", queryParams=queryParams)
  return to[seq[address]](resp.body)

