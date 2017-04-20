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


proc getUsers*(srv: Users_service, queryParams: Table[string, string] = initTable[string, string]()) : usersGetRespBody =
  # get users.# This method will be return list user.# Use it wisely.
  # It calls GET /users endpoint.

  let resp = srv.client.request("/users", "GET", queryParams=queryParams)
  return to[usersGetRespBody](resp.body)

proc usersPost*(srv: Users_service, reqBody: City, queryParams: Table[string, string] = initTable[string, string]()) : City =
  # create users
  # It calls POST /users endpoint.

  let resp = srv.client.request("/users", "POST", $$reqBody, queryParams=queryParams)
  return to[City](resp.body)

proc usersPut*(srv: Users_service, reqBody: seq[string], queryParams: Table[string, string] = initTable[string, string]()) : City =
  # create users
  # It calls PUT /users endpoint.

  let resp = srv.client.request("/users", "PUT", $$reqBody, queryParams=queryParams)
  return to[City](resp.body)

proc OptionsUsers*(srv: Users_service, queryParams: Table[string, string] = initTable[string, string]()) : string =
  
  # It calls OPTIONS /users endpoint.

  let resp = srv.client.request("/users", "OPTIONS", queryParams=queryParams)
  return to[string](resp.body)

proc GetUserByID*(srv: Users_service, userId: string, queryParams: Table[string, string] = initTable[string, string]()) : City =
  # get id
  # It calls GET /users/{userId} endpoint.

  let resp = srv.client.request("/users/"&userId, "GET", queryParams=queryParams)
  return to[City](resp.body)

proc usersByUserIdDelete*(srv: Users_service, userId: string, queryParams: Table[string, string] = initTable[string, string]()) : string =
  
  # It calls DELETE /users/{userId} endpoint.

  let resp = srv.client.request("/users/"&userId, "DELETE", queryParams=queryParams)
  return to[string](resp.body)

proc usersByUserIdAddressPost*(srv: Users_service, reqBody: usersuserIdaddressPostReqBody, userId: string, queryParams: Table[string, string] = initTable[string, string]()) : usersuserIdaddressPostRespBody =
  
  # It calls POST /users/{userId}/address endpoint.

  let resp = srv.client.request("/users/"&userId&"/address", "POST", $$reqBody, queryParams=queryParams)
  return to[usersuserIdaddressPostRespBody](resp.body)

proc usersByUserIdAddressFolderAddressIdTestAddressId2Get*(srv: Users_service, addressId: string, addressId2: string, userId: string, queryParams: Table[string, string] = initTable[string, string]()) : seq[address] =
  # get address id
  # It calls GET /users/{userId}/address/folder{addressId}test{addressId2} endpoint.

  let resp = srv.client.request("/users/"&userId&"/address/folder"&addressId&"test"&addressId2, "GET", queryParams=queryParams)
  return to[seq[address]](resp.body)

