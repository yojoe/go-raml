import jester, marshal

import User


proc usersGet*(req: Request) : tuple[code: HttpCode, content: User] =
  # Get random user
  
  var respBody: User
  respBody = User(name:"John", username: "Doe")
  result = (code: Http200, content: respBody)

proc usersPost*(req: Request) : tuple[code: HttpCode, content: string] =
  # Add user
  let reqBody = to[User](req.body)
  var respBody: string
  respBody = reqBody.name
  result = (code: Http200, content: respBody)

proc usersByUsernameGet*(username: string, req: Request) : tuple[code: HttpCode, content: User] =
  # Get information on a specific user
  
  var respBody: User
  result = (code: Http200, content: respBody)

