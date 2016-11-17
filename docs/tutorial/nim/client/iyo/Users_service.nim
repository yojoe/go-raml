import marshal
import client_itsyouonline

import Address
import Authorization
import BankAccount
import Contract
import DigitalAssetAddress
import JoinOrganizationInvitation
import Phonenumber
import PublicKey
import RegistryEntry
import User
import UserAPIKey
import userview


type
  Users_service* = object
    client*: Client
    name*: string

proc UsersSrv*(c : Client) : Users_service  =
  return Users_service(client:c, name:c.baseURI)


proc CreateUser*(srv: Users_service, reqBody: User) : string =
  let resp = srv.client.request("/users", "POST", $$reqBody)
  return to[string](resp.body)

proc GetUserInformation*(srv: Users_service, username: string) : userview =
  let resp = srv.client.request("/users/"&username&"/info", "GET")
  return to[userview](resp.body)

proc ValidateUsername*(srv: Users_service, username: string) : string =
  let resp = srv.client.request("/users/"&username&"/validate", "GET")
  return to[string](resp.body)

proc GetUserAddresses*(srv: Users_service, username: string) : seq[Address] =
  let resp = srv.client.request("/users/"&username&"/addresses", "GET")
  return to[seq[Address]](resp.body)

proc RegisterNewUserAddress*(srv: Users_service, reqBody: Address, username: string) : Address =
  let resp = srv.client.request("/users/"&username&"/addresses", "POST", $$reqBody)
  return to[Address](resp.body)

proc GetUserAddressByLabel*(srv: Users_service, label: string, username: string) : Address =
  let resp = srv.client.request("/users/"&username&"/addresses/"&label, "GET")
  return to[Address](resp.body)

proc UpdateUserAddress*(srv: Users_service, reqBody: Address, label: string, username: string) : string =
  let resp = srv.client.request("/users/"&username&"/addresses/"&label, "PUT", $$reqBody)
  return to[string](resp.body)

proc DeleteUserAddress*(srv: Users_service, label: string, username: string) : string =
  let resp = srv.client.request("/users/"&username&"/addresses/"&label, "DELETE")
  return to[string](resp.body)

proc GetDigitalWallet*(srv: Users_service, username: string) : seq[DigitalAssetAddress] =
  let resp = srv.client.request("/users/"&username&"/digitalwallet", "GET")
  return to[seq[DigitalAssetAddress]](resp.body)

proc RegisterNewDigitalAssetAddress*(srv: Users_service, reqBody: DigitalAssetAddress, username: string) : DigitalAssetAddress =
  let resp = srv.client.request("/users/"&username&"/digitalwallet", "POST", $$reqBody)
  return to[DigitalAssetAddress](resp.body)

proc GetDigitalAssetAddressByLabel*(srv: Users_service, label: string, username: string) : DigitalAssetAddress =
  let resp = srv.client.request("/users/"&username&"/digitalwallet/"&label, "GET")
  return to[DigitalAssetAddress](resp.body)

proc UpdateDigitalAssetAddress*(srv: Users_service, reqBody: DigitalAssetAddress, label: string, username: string) : string =
  let resp = srv.client.request("/users/"&username&"/digitalwallet/"&label, "PUT", $$reqBody)
  return to[string](resp.body)

proc DeleteDigitalAssetAddress*(srv: Users_service, label: string, username: string) : string =
  let resp = srv.client.request("/users/"&username&"/digitalwallet/"&label, "DELETE")
  return to[string](resp.body)

proc GetUserPhoneNumbers*(srv: Users_service, username: string) : seq[Phonenumber] =
  let resp = srv.client.request("/users/"&username&"/phonenumbers", "GET")
  return to[seq[Phonenumber]](resp.body)

proc RegisterNewUserPhonenumber*(srv: Users_service, reqBody: Phonenumber, username: string) : Phonenumber =
  let resp = srv.client.request("/users/"&username&"/phonenumbers", "POST", $$reqBody)
  return to[Phonenumber](resp.body)

proc GetUserPhonenumberByLabel*(srv: Users_service, label: string, username: string) : Phonenumber =
  let resp = srv.client.request("/users/"&username&"/phonenumbers/"&label, "GET")
  return to[Phonenumber](resp.body)

proc UpdateUserPhonenumber*(srv: Users_service, reqBody: Phonenumber, label: string, username: string) : string =
  let resp = srv.client.request("/users/"&username&"/phonenumbers/"&label, "PUT", $$reqBody)
  return to[string](resp.body)

proc DeleteUserPhonenumber*(srv: Users_service, label: string, username: string) : string =
  let resp = srv.client.request("/users/"&username&"/phonenumbers/"&label, "DELETE")
  return to[string](resp.body)

proc ValidatePhonenumber*(srv: Users_service, label: string, username: string) : usersusernamephonenumberslabelactivatePostRespBody =
  let resp = srv.client.request("/users/"&username&"/phonenumbers/"&label&"/activate", "POST")
  return to[usersusernamephonenumberslabelactivatePostRespBody](resp.body)

proc VerifyPhoneNumber*(srv: Users_service, reqBody: usersusernamephonenumberslabelactivatePutReqBody, label: string, username: string) : string =
  let resp = srv.client.request("/users/"&username&"/phonenumbers/"&label&"/activate", "PUT", $$reqBody)
  return to[string](resp.body)

proc GetUserBankAccounts*(srv: Users_service, username: string) : seq[BankAccount] =
  let resp = srv.client.request("/users/"&username&"/banks", "GET")
  return to[seq[BankAccount]](resp.body)

proc CreateUserBankAccount*(srv: Users_service, reqBody: usersusernamebanksPostReqBody, username: string) : usersusernamebanksPostRespBody =
  let resp = srv.client.request("/users/"&username&"/banks", "POST", $$reqBody)
  return to[usersusernamebanksPostRespBody](resp.body)

proc GetUserBankAccountByLabel*(srv: Users_service, username: string, label: string) : BankAccount =
  let resp = srv.client.request("/users/"&username&"/banks/"&label, "GET")
  return to[BankAccount](resp.body)

proc UpdateUserBankAccount*(srv: Users_service, reqBody: BankAccount, username: string, label: string) : string =
  let resp = srv.client.request("/users/"&username&"/banks/"&label, "PUT", $$reqBody)
  return to[string](resp.body)

proc DeleteUserBankAccount*(srv: Users_service, username: string, label: string) : string =
  let resp = srv.client.request("/users/"&username&"/banks/"&label, "DELETE")
  return to[string](resp.body)

proc GetUser*(srv: Users_service, username: string) : User =
  let resp = srv.client.request("/users/"&username, "GET")
  return to[User](resp.body)

proc UpdatePassword*(srv: Users_service, reqBody: usersusernamepasswordPutReqBody, username: string) : string =
  let resp = srv.client.request("/users/"&username&"/password", "PUT", $$reqBody)
  return to[string](resp.body)

proc GetEmailAddresses*(srv: Users_service, username: string) : seq[EmailAddress] =
  let resp = srv.client.request("/users/"&username&"/emailaddresses", "GET")
  return to[seq[EmailAddress]](resp.body)

proc RegisterNewEmailAddress*(srv: Users_service, reqBody: usersusernameemailaddressesPostReqBody, username: string) : usersusernameemailaddressesPostRespBody =
  let resp = srv.client.request("/users/"&username&"/emailaddresses", "POST", $$reqBody)
  return to[usersusernameemailaddressesPostRespBody](resp.body)

proc UpdateEmailAddress*(srv: Users_service, reqBody: usersusernameemailaddresseslabelPutReqBody, label: string, username: string) : string =
  let resp = srv.client.request("/users/"&username&"/emailaddresses/"&label, "PUT", $$reqBody)
  return to[string](resp.body)

proc DeleteEmailAddress*(srv: Users_service, label: string, username: string) : string =
  let resp = srv.client.request("/users/"&username&"/emailaddresses/"&label, "DELETE")
  return to[string](resp.body)

proc ValidateEmailAddress*(srv: Users_service, label: string, username: string) : string =
  let resp = srv.client.request("/users/"&username&"/emailaddresses/"&label&"/validate", "POST")
  return to[string](resp.body)

proc ListAPIKeys*(srv: Users_service, username: string) : seq[UserAPIKey] =
  let resp = srv.client.request("/users/"&username&"/apikeys", "GET")
  return to[seq[UserAPIKey]](resp.body)

proc AddApiKey*(srv: Users_service, reqBody: usersusernameapikeysPostReqBody, username: string) : UserAPIKey =
  let resp = srv.client.request("/users/"&username&"/apikeys", "POST", $$reqBody)
  return to[UserAPIKey](resp.body)

proc GetAPIkey*(srv: Users_service, label: string, username: string) : UserAPIKey =
  let resp = srv.client.request("/users/"&username&"/apikeys/"&label, "GET")
  return to[UserAPIKey](resp.body)

proc UpdateAPIkey*(srv: Users_service, reqBody: usersusernameapikeyslabelPutReqBody, label: string, username: string) : string =
  let resp = srv.client.request("/users/"&username&"/apikeys/"&label, "PUT", $$reqBody)
  return to[string](resp.body)

proc DeleteAPIkey*(srv: Users_service, label: string, username: string) : string =
  let resp = srv.client.request("/users/"&username&"/apikeys/"&label, "DELETE")
  return to[string](resp.body)

proc ListUserRegistry*(srv: Users_service, username: string) : seq[RegistryEntry] =
  let resp = srv.client.request("/users/"&username&"/registry", "GET")
  return to[seq[RegistryEntry]](resp.body)

proc AddUserRegistryEntry*(srv: Users_service, reqBody: RegistryEntry, username: string) : RegistryEntry =
  let resp = srv.client.request("/users/"&username&"/registry", "POST", $$reqBody)
  return to[RegistryEntry](resp.body)

proc GetUserRegistryEntry*(srv: Users_service, key: string, username: string) : RegistryEntry =
  let resp = srv.client.request("/users/"&username&"/registry/"&key, "GET")
  return to[RegistryEntry](resp.body)

proc DeleteUserRegistryEntry*(srv: Users_service, key: string, username: string) : string =
  let resp = srv.client.request("/users/"&username&"/registry/"&key, "DELETE")
  return to[string](resp.body)

proc DeleteGithubAccount*(srv: Users_service, username: string) : string =
  let resp = srv.client.request("/users/"&username&"/github", "DELETE")
  return to[string](resp.body)

proc GetTwoFAMethods*(srv: Users_service, username: string) : usersusernametwofamethodsGetRespBody =
  let resp = srv.client.request("/users/"&username&"/twofamethods", "GET")
  return to[usersusernametwofamethodsGetRespBody](resp.body)

proc UpdateUserName*(srv: Users_service, reqBody: usersusernamenamePutReqBody, username: string) : string =
  let resp = srv.client.request("/users/"&username&"/name", "PUT", $$reqBody)
  return to[string](resp.body)

proc GetTOTPSecret*(srv: Users_service, username: string) : usersusernametotpGetRespBody =
  let resp = srv.client.request("/users/"&username&"/totp", "GET")
  return to[usersusernametotpGetRespBody](resp.body)

proc SetupTOTP*(srv: Users_service, reqBody: usersusernametotpPostReqBody, username: string) : string =
  let resp = srv.client.request("/users/"&username&"/totp", "POST", $$reqBody)
  return to[string](resp.body)

proc RemoveTOTP*(srv: Users_service, username: string) : string =
  let resp = srv.client.request("/users/"&username&"/totp", "DELETE")
  return to[string](resp.body)

proc DeleteFacebookAccount*(srv: Users_service, username: string) : string =
  let resp = srv.client.request("/users/"&username&"/facebook", "DELETE")
  return to[string](resp.body)

proc GetUserOrganizations*(srv: Users_service, username: string) : usersusernameorganizationsGetRespBody =
  let resp = srv.client.request("/users/"&username&"/organizations", "GET")
  return to[usersusernameorganizationsGetRespBody](resp.body)

proc LeaveOrganization*(srv: Users_service, globalid: string, username: string) : string =
  let resp = srv.client.request("/users/"&username&"/organizations/"&globalid&"/leave", "DELETE")
  return to[string](resp.body)

proc AcceptMembership*(srv: Users_service, reqBody: JoinOrganizationInvitation, role: string, globalid: string, username: string) : JoinOrganizationInvitation =
  let resp = srv.client.request("/users/"&username&"/organizations/"&globalid&"/roles/"&role, "POST", $$reqBody)
  return to[JoinOrganizationInvitation](resp.body)

proc usersByUsernameOrganizationsByGlobalidRolesByRoleDelete*(srv: Users_service, role: string, globalid: string, username: string) : string =
  let resp = srv.client.request("/users/"&username&"/organizations/"&globalid&"/roles/"&role, "DELETE")
  return to[string](resp.body)

proc ListPublicKeys*(srv: Users_service, username: string) : seq[PublicKey] =
  let resp = srv.client.request("/users/"&username&"/publickeys", "GET")
  return to[seq[PublicKey]](resp.body)

proc AddPublicKey*(srv: Users_service, reqBody: PublicKey, username: string) : PublicKey =
  let resp = srv.client.request("/users/"&username&"/publickeys", "POST", $$reqBody)
  return to[PublicKey](resp.body)

proc GetPublicKey*(srv: Users_service, label: string, username: string) : PublicKey =
  let resp = srv.client.request("/users/"&username&"/publickeys/"&label, "GET")
  return to[PublicKey](resp.body)

proc UpdatePublicKey*(srv: Users_service, reqBody: PublicKey, label: string, username: string) : PublicKey =
  let resp = srv.client.request("/users/"&username&"/publickeys/"&label, "PUT", $$reqBody)
  return to[PublicKey](resp.body)

proc DeletePublicKey*(srv: Users_service, label: string, username: string) : string =
  let resp = srv.client.request("/users/"&username&"/publickeys/"&label, "DELETE")
  return to[string](resp.body)

proc GetAllAuthorizations*(srv: Users_service, username: string) : seq[Authorization] =
  let resp = srv.client.request("/users/"&username&"/authorizations", "GET")
  return to[seq[Authorization]](resp.body)

proc GetAuthorization*(srv: Users_service, grantedTo: string, username: string) : Authorization =
  let resp = srv.client.request("/users/"&username&"/authorizations/"&grantedTo, "GET")
  return to[Authorization](resp.body)

proc UpdateAuthorization*(srv: Users_service, reqBody: Authorization, grantedTo: string, username: string) : string =
  let resp = srv.client.request("/users/"&username&"/authorizations/"&grantedTo, "PUT", $$reqBody)
  return to[string](resp.body)

proc DeleteAuthorization*(srv: Users_service, grantedTo: string, username: string) : string =
  let resp = srv.client.request("/users/"&username&"/authorizations/"&grantedTo, "DELETE")
  return to[string](resp.body)

proc GetUserContracts*(srv: Users_service, username: string) : string =
  let resp = srv.client.request("/users/"&username&"/contracts", "GET")
  return to[string](resp.body)

proc CreateUserContract*(srv: Users_service, reqBody: Contract, username: string) : Contract =
  let resp = srv.client.request("/users/"&username&"/contracts", "POST", $$reqBody)
  return to[Contract](resp.body)

proc GetNotifications*(srv: Users_service, username: string) : usersusernamenotificationsGetRespBody =
  let resp = srv.client.request("/users/"&username&"/notifications", "GET")
  return to[usersusernamenotificationsGetRespBody](resp.body)

