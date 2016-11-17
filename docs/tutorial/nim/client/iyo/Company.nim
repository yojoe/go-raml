
import times
type
  Company* = object
    expire*: Time
    globalid*: string
    info*: seq[string]
    organizations*: seq[string]
    publicKeys*: seq[string]
    taxnr*: string
