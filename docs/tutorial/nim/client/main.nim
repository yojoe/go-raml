import ./goraml/client_goraml, ./goraml/Users_service, ./goraml/oauth2_client_itsyouonline

let clientId = ""
let clientSecret = ""

# create the client
let c = client_goraml.newClient()

# Example of calling endpoint which doesn't need scope
# get JWT token from itsyou.online and then set it as authorization header
let jwtToken = c.getAccessToken(clientSecret, clientId, @[], @[])
c.setAuthHeader("Bearer " & jwtToken)

let resp = c.UsersSrv.usersGet()
echo "resp=", $resp

# Example of calling endpoint which need scope
# get JWT token from itsyou.online and set it as authorization header
#let jwtToken = c.getTokenByClientCrendentials(clientId, clientSecret, @["user:memberof:goraml"], @["external1"])
#c.setAuthHeader("token " & jwtToken)
#let resp2 = c.UsersSrv.usersByUsernameGet("john")
#echo "resp2=", $resp2
