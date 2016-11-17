import ./iyo/client_itsyouonline as iyoClient
import iyo/oauth2_client_itsyouonline
import ./goraml/client_goraml, ./goraml/Users_service

# get JWT token from itsyou.online
let iyo = iyoClient.newClient()

let token = iyo.getAccessToken("client-id", "client-secret")
echo "token=", token

let jwtToken = iyo.createJWTToken(@["user:memberof:goraml"], @["external1"])

# make request to goraml server
let c = client_goraml.newClient()
c.setAuthHeader("token " & jwtToken)

let resp = c.UsersSrv.usersGet()
echo "resp=", $resp
