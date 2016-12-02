import jester, asyncdispatch, json, marshal, system

import users_api

routes:
  GET "/users":
    let ret = usersGet(request)
    resp(ret.code, $$ret.content)

  POST "/users":
    let ret = usersPost(request)
    resp(ret.code, $$ret.content)

  GET "/users/@username":
    let ret = usersByUsernameGet(@"username", request)
    resp(ret.code, $$ret.content)


  GET "/":
    #resp(readFile("index.html"))
    resp("Hello World!")

runForever()
