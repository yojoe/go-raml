
type
  Authorization* = object
    addresses*: string
    bankaccounts*: string
    emailaddresses*: string
    facebook*: bool
    github*: bool
    grantedTo*: string
    organizations*: seq[string]
    phonenumbers*: string
    publicKeys*: seq[AuthorizationMap]
    username*: string
