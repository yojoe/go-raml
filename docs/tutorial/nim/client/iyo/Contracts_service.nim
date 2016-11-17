import marshal
import client_itsyouonline

import Contract
import Signature


type
  Contracts_service* = object
    client*: Client
    name*: string

proc ContractsSrv*(c : Client) : Contracts_service  =
  return Contracts_service(client:c, name:c.baseURI)


proc GetContract*(srv: Contracts_service, contractId: string) : Contract =
  let resp = srv.client.request("/contracts/"&contractId, "GET")
  return to[Contract](resp.body)

proc SignContract*(srv: Contracts_service, reqBody: Signature, contractId: string) : string =
  let resp = srv.client.request("/contracts/"&contractId&"/signatures", "POST", $$reqBody)
  return to[string](resp.body)

