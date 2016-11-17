
import times
type
  DigitalAssetAddress* = object
    address*: string
    currencysymbol*: string
    expire*: Time
    label*: string
    noexpiration*: bool
