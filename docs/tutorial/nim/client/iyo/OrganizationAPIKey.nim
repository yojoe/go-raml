
import Label
type
  OrganizationAPIKey* = object
    callbackURL*: string
    clientCredentialsGrantType*: bool
    label*: Label
    secret*: string
