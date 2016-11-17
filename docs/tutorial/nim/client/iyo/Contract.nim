
import times
type
  Contract* = object
    content*: string
    contractId*: string
    contractType*: string
    expires*: Time
    extends*: seq[string]
    invalidates*: seq[string]
    parties*: seq[Party]
    signatures*: seq[Signature]
