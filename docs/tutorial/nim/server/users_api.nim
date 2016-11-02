import jester, marshal

import User


proc usersGet*(req: Request) : tuple[code: HttpCode, content: seq[User]] =
  # Get list of all developers
  
  var respBody: seq[User]
  let user1 = User(name:"john", username:"doe") ## added line
  respBody = @[user1] #added line
  result = (code: Http200, content: respBody)

proc usersPost*(req: Request) : tuple[code: HttpCode, content: User] =
  # Add user
  let reqBody = to[User](req.body)
  var respBody: User
  respBody = User(name:reqBody.name, username: reqBody.username) #added line
  result = (code: Http200, content: respBody)

proc usersByUsernameGet*(username: string, req: Request) : tuple[code: HttpCode, content: User] =
  # Get information on a specific user
  
  var respBody: User
  respBody = User(name: username, username: username) #added line
  result = (code: Http200, content: respBody)

