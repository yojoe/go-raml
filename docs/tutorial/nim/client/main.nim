import Users_service, client, User

let c = newClient()

let u = c.UsersSrv.usersByUsernameGet("paijo")
echo "name=",u.name

let us = c.UsersSrv.usersGet()
let u1 = us[0]
echo "u1name=",u1.name

let up = User(name:"john", username:"doe")
let up_res = c.UsersSrv.usersPost(up)
echo "up_res.name=",up_res.name, ".up_res.username=", up_res.username

