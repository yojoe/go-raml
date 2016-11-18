import ./goraml/client_goraml, ./goraml/Users_service, ./goraml/oauth2_client_itsyouonline

# create the client
let c = client_goraml.newClient()

# get JWT token from itsyou.online
let jwtToken = c.getTokenByClientCrendentials("client-id", "client-secret", @["user:memberof:goraml"], @["external1"])

# set token as Authorization header
echo "jwt token=", jwtToken
c.setAuthHeader("token " & jwtToken)

# make request to goraml server
let resp = c.UsersSrv.usersGet()
echo "resp=", $resp
