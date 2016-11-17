
import FacebookAccount
import GithubAccount
type
  userview* = object
    addresses*: seq[Address]
    bankaccounts*: seq[BankAccount]
    emailaddresses*: seq[EmailAddress]
    facebook*: FacebookAccount
    github*: GithubAccount
    organizations*: seq[string]
    phonenumbers*: seq[Phonenumber]
    publicKeys*: seq[PublicKey]
    username*: string
