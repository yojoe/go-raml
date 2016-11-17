
type
  Organization* = object
    dns*: seq[string]
    globalid*: string
    includes*: seq[string]
    members*: seq[string]
    orgmembers*: seq[string]
    orgowners*: seq[string]
    owners*: seq[string]
    publicKeys*: seq[string]
