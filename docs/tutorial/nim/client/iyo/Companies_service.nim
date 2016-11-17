import marshal
import client_itsyouonline

import Company
import Contract
import companyview


type
  Companies_service* = object
    client*: Client
    name*: string

proc CompaniesSrv*(c : Client) : Companies_service  =
  return Companies_service(client:c, name:c.baseURI)


proc GetCompanyList*(srv: Companies_service) : seq[Company] =
  let resp = srv.client.request("/companies", "GET")
  return to[seq[Company]](resp.body)

proc CreateCompany*(srv: Companies_service, reqBody: Company) : string =
  let resp = srv.client.request("/companies", "POST", $$reqBody)
  return to[string](resp.body)

proc GetCompany*(srv: Companies_service, globalId: string) : Company =
  let resp = srv.client.request("/companies/"&globalId, "GET")
  return to[Company](resp.body)

proc UpdateCompany*(srv: Companies_service, globalId: string) : Company =
  let resp = srv.client.request("/companies/"&globalId, "PUT")
  return to[Company](resp.body)

proc GetCompanyContracts*(srv: Companies_service, globalId: string) : string =
  let resp = srv.client.request("/companies/"&globalId&"/contracts", "GET")
  return to[string](resp.body)

proc CreateCompanyContract*(srv: Companies_service, reqBody: Contract, globalId: string) : Contract =
  let resp = srv.client.request("/companies/"&globalId&"/contracts", "POST", $$reqBody)
  return to[Contract](resp.body)

proc GetCompanyInfo*(srv: Companies_service, globalId: string) : companyview =
  let resp = srv.client.request("/companies/"&globalId&"/info", "GET")
  return to[companyview](resp.body)

proc companiesByGlobalIdValidateGet*(srv: Companies_service, globalId: string) : string =
  let resp = srv.client.request("/companies/"&globalId&"/validate", "GET")
  return to[string](resp.body)

