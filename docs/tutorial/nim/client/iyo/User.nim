
import FacebookAccount
import GithubAccount
import times
type
  User* = object
    addresses*: seq[Address]
    bankaccounts*: seq[BankAccount]
    digitalwallet*: seq[DigitalAssetAddress]
    emailaddresses*: seq[EmailAddress]
    expire*: Time
    facebook*: FacebookAccount
    firstname*: string
    github*: GithubAccount
    lastname*: string
    phonenumbers*: seq[Phonenumber]
    publicKeys*: seq[string]
    username*: string
