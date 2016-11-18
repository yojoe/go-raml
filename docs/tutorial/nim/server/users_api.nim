import jester, marshal, system
import oauth2_jwt

import User

let ojwt = Oauth2JWT(pubKey:readFile("itsyouonline.pub"))


proc usersGet*(req: Request) : tuple[code: HttpCode, content: seq[User]] =
  # Get list of all developers
  var respBody: seq[User]
  if not ojwt.checkJWTToken(req, @[]): return (code: Http403, content: respBody)
  
  result = (code: Http200, content: respBody)

proc usersPost*(req: Request) : tuple[code: HttpCode, content: User] =
  # Add user
  var respBody: User
  if not ojwt.checkJWTToken(req, @["user:memberof:goraml-admin"]): return (code: Http403, content: respBody)
  let reqBody = to[User](req.body)
  result = (code: Http200, content: respBody)

proc usersByUsernameGet*(username: string, req: Request) : tuple[code: HttpCode, content: User] =
  # Get information on a specific user
  var respBody: User
  if not ojwt.checkJWTToken(req, @["user:memberof:goraml"]): return (code: Http403, content: respBody)
  
  result = (code: Http200, content: respBody)

