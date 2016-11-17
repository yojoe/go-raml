
import Label
type
  UserAPIKey* = object
    apikey*: string
    applicationid*: string
    label*: Label
    scopes*: seq[string]
    username*: string
