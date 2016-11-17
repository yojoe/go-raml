
type
  GetNotificationsRespBody* = object
    approvals*: seq[JoinOrganizationInvitation]
    contractRequests*: seq[ContractSigningRequest]
    invitations*: seq[JoinOrganizationInvitation]
