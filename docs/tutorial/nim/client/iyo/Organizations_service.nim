import marshal
import client_itsyouonline

import Contract
import Member
import Organization
import OrganizationAPIKey
import RegistryEntry


type
  Organizations_service* = object
    client*: Client
    name*: string

proc OrganizationsSrv*(c : Client) : Organizations_service  =
  return Organizations_service(client:c, name:c.baseURI)


proc CreateNewOrganization*(srv: Organizations_service, reqBody: Organization) : Organization =
  let resp = srv.client.request("/organizations", "POST", $$reqBody)
  return to[Organization](resp.body)

proc GetOrganization*(srv: Organizations_service, globalid: string) : Organization =
  let resp = srv.client.request("/organizations/"&globalid, "GET")
  return to[Organization](resp.body)

proc CreateNewSubOrganization*(srv: Organizations_service, reqBody: Organization, globalid: string) : Organization =
  let resp = srv.client.request("/organizations/"&globalid, "POST", $$reqBody)
  return to[Organization](resp.body)

proc UpdateOrganization*(srv: Organizations_service, reqBody: Organization, globalid: string) : Organization =
  let resp = srv.client.request("/organizations/"&globalid, "PUT", $$reqBody)
  return to[Organization](resp.body)

proc DeleteOrganization*(srv: Organizations_service, globalid: string) : string =
  let resp = srv.client.request("/organizations/"&globalid, "DELETE")
  return to[string](resp.body)

proc CreateOrganizationDNS*(srv: Organizations_service, reqBody: organizationsglobaliddnsdnsnamePostReqBody, dnsname: string, globalid: string) : organizationsglobaliddnsdnsnamePostRespBody =
  let resp = srv.client.request("/organizations/"&globalid&"/dns/"&dnsname, "POST", $$reqBody)
  return to[organizationsglobaliddnsdnsnamePostRespBody](resp.body)

proc UpdateOrganizationDNS*(srv: Organizations_service, reqBody: organizationsglobaliddnsdnsnamePutReqBody, dnsname: string, globalid: string) : string =
  let resp = srv.client.request("/organizations/"&globalid&"/dns/"&dnsname, "PUT", $$reqBody)
  return to[string](resp.body)

proc DeleteOrganizaitonDNS*(srv: Organizations_service, dnsname: string, globalid: string) : string =
  let resp = srv.client.request("/organizations/"&globalid&"/dns/"&dnsname, "DELETE")
  return to[string](resp.body)

proc Get2faValidityTime*(srv: Organizations_service, globalid: string) : int =
  let resp = srv.client.request("/organizations/"&globalid&"/2fa/validity", "GET")
  return to[int](resp.body)

proc Set2faValidityTime*(srv: Organizations_service, reqBody: int, globalid: string) : string =
  let resp = srv.client.request("/organizations/"&globalid&"/2fa/validity", "POST", $$reqBody)
  return to[string](resp.body)

proc SetOrgOwner*(srv: Organizations_service, reqBody: organizationsglobalidorgownersPostReqBody, globalid: string) : string =
  let resp = srv.client.request("/organizations/"&globalid&"/orgowners", "POST", $$reqBody)
  return to[string](resp.body)

proc DeleteOrgOwner*(srv: Organizations_service, globalid2: string, globalid: string) : string =
  let resp = srv.client.request("/organizations/"&globalid&"/orgowners/"&globalid2, "DELETE")
  return to[string](resp.body)

proc AddOrganizationMember*(srv: Organizations_service, reqBody: Member, globalid: string) : Member =
  let resp = srv.client.request("/organizations/"&globalid&"/members", "POST", $$reqBody)
  return to[Member](resp.body)

proc UpdateOrganizationMemberShip*(srv: Organizations_service, reqBody: organizationsglobalidmembersPutReqBody, globalid: string) : Organization =
  let resp = srv.client.request("/organizations/"&globalid&"/members", "PUT", $$reqBody)
  return to[Organization](resp.body)

proc RemoveOrganizationMember*(srv: Organizations_service, username: string, globalid: string) : string =
  let resp = srv.client.request("/organizations/"&globalid&"/members/"&username, "DELETE")
  return to[string](resp.body)

proc AddOrganizationOwner*(srv: Organizations_service, reqBody: Member, globalid: string) : Member =
  let resp = srv.client.request("/organizations/"&globalid&"/owners", "POST", $$reqBody)
  return to[Member](resp.body)

proc RemoveOrganizationOwner*(srv: Organizations_service, username: string, globalid: string) : string =
  let resp = srv.client.request("/organizations/"&globalid&"/owners/"&username, "DELETE")
  return to[string](resp.body)

proc GetOrganizationContracts*(srv: Organizations_service, globalid: string) : string =
  let resp = srv.client.request("/organizations/"&globalid&"/contracts", "GET")
  return to[string](resp.body)

proc CreateOrganizationContracty*(srv: Organizations_service, reqBody: Contract, globalid: string) : Contract =
  let resp = srv.client.request("/organizations/"&globalid&"/contracts", "POST", $$reqBody)
  return to[Contract](resp.body)

proc GetPendingOrganizationInvitations*(srv: Organizations_service, globalid: string) : seq[JoinOrganizationInvitation] =
  let resp = srv.client.request("/organizations/"&globalid&"/invitations", "GET")
  return to[seq[JoinOrganizationInvitation]](resp.body)

proc RemovePendingOrganizationInvitation*(srv: Organizations_service, username: string, globalid: string) : string =
  let resp = srv.client.request("/organizations/"&globalid&"/invitations/"&username, "DELETE")
  return to[string](resp.body)

proc ListOrganizationRegistry*(srv: Organizations_service, globalid: string) : seq[RegistryEntry] =
  let resp = srv.client.request("/organizations/"&globalid&"/registry", "GET")
  return to[seq[RegistryEntry]](resp.body)

proc AddOrganizationRegistryEntry*(srv: Organizations_service, reqBody: RegistryEntry, globalid: string) : RegistryEntry =
  let resp = srv.client.request("/organizations/"&globalid&"/registry", "POST", $$reqBody)
  return to[RegistryEntry](resp.body)

proc GetOrganizationTree*(srv: Organizations_service, globalid: string) : seq[OrganizationTreeItem] =
  let resp = srv.client.request("/organizations/"&globalid&"/tree", "GET")
  return to[seq[OrganizationTreeItem]](resp.body)

proc GetOrganizationLogo*(srv: Organizations_service, globalid: string) : string =
  let resp = srv.client.request("/organizations/"&globalid&"/logo", "GET")
  return to[string](resp.body)

proc SetOrganizationLogo*(srv: Organizations_service, reqBody: organizationsglobalidlogoPutReqBody, globalid: string) : string =
  let resp = srv.client.request("/organizations/"&globalid&"/logo", "PUT", $$reqBody)
  return to[string](resp.body)

proc DeleteOrganizationLogo*(srv: Organizations_service, globalid: string) : string =
  let resp = srv.client.request("/organizations/"&globalid&"/logo", "DELETE")
  return to[string](resp.body)

proc SetOrgMember*(srv: Organizations_service, reqBody: organizationsglobalidorgmembersPostReqBody, globalid: string) : string =
  let resp = srv.client.request("/organizations/"&globalid&"/orgmembers", "POST", $$reqBody)
  return to[string](resp.body)

proc UpdateOrganizationOrgMemberShip*(srv: Organizations_service, reqBody: organizationsglobalidorgmembersPutReqBody, globalid: string) : organizationsglobalidorgmembersPutRespBody =
  let resp = srv.client.request("/organizations/"&globalid&"/orgmembers", "PUT", $$reqBody)
  return to[organizationsglobalidorgmembersPutRespBody](resp.body)

proc DeleteOrgMember*(srv: Organizations_service, globalid2: string, globalid: string) : string =
  let resp = srv.client.request("/organizations/"&globalid&"/orgmembers/"&globalid2, "DELETE")
  return to[string](resp.body)

proc GetOrganizationAPIKeyLabels*(srv: Organizations_service, globalid: string) : seq[string] =
  let resp = srv.client.request("/organizations/"&globalid&"/apikeys", "GET")
  return to[seq[string]](resp.body)

proc CreateNewOrganizationAPIKey*(srv: Organizations_service, reqBody: OrganizationAPIKey, globalid: string) : OrganizationAPIKey =
  let resp = srv.client.request("/organizations/"&globalid&"/apikeys", "POST", $$reqBody)
  return to[OrganizationAPIKey](resp.body)

proc GetOrganizationAPIKey*(srv: Organizations_service, label: string, globalid: string) : OrganizationAPIKey =
  let resp = srv.client.request("/organizations/"&globalid&"/apikeys/"&label, "GET")
  return to[OrganizationAPIKey](resp.body)

proc UpdateOrganizationAPIKey*(srv: Organizations_service, reqBody: organizationsglobalidapikeyslabelPutReqBody, label: string, globalid: string) : string =
  let resp = srv.client.request("/organizations/"&globalid&"/apikeys/"&label, "PUT", $$reqBody)
  return to[string](resp.body)

proc DeleteOrganizationAPIKey*(srv: Organizations_service, label: string, globalid: string) : string =
  let resp = srv.client.request("/organizations/"&globalid&"/apikeys/"&label, "DELETE")
  return to[string](resp.body)

proc GetOrganizationRegistryEntry*(srv: Organizations_service, key: string, globalid: string) : RegistryEntry =
  let resp = srv.client.request("/organizations/"&globalid&"/registry/"&key, "GET")
  return to[RegistryEntry](resp.body)

proc DeleteOrganizationRegistryEntry*(srv: Organizations_service, key: string, globalid: string) : string =
  let resp = srv.client.request("/organizations/"&globalid&"/registry/"&key, "DELETE")
  return to[string](resp.body)

