
type
  GetTwoFAMethodsRespBody* = object
    sms*: seq[Phonenumber]
    totp*: bool
