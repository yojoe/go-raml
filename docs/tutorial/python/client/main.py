import sys

from client import Client
import itsyouonline

# itsyou.online client
iyo_client = itsyouonline.Client()

# goramldir client
client = Client()

def main(app_id, app_secret):
    '''
    Login to itsyou.online server
    If succeed, next request will use `Authorization` header
    acquired from this login process
    '''
    iyo_client.oauth.LoginViaClientCredentials(app_id, app_secret)

    '''
    create JWT token with specified 'scopes'.
    You need to change 'user:memberof:goraml' to match with your
    organization scopes
    '''
    jwt_token = iyo_client.oauth.CreateJWTToken(["user:memberof:goraml"])

    # Set our goramldir client to use JWT token from itsyou.online
    client.set_auth_header("token " + jwt_token)

    # try to make simple GET call to goramldir server
    resp = client.users_byUsername_get("john")
    print(resp.json())

if __name__ == "__main__":
    '''
    usage : python3 main.py application_id application_secret
    '''
    main(sys.argv[1], sys.argv[2])
