import sys

import goramldir

# goramldir client
client = goramldir.Client()

def main(app_id, app_secret):
    # get JWT token from itsyouonline
    jwt_token = client.oauth2_client_itsyouonline.get_access_token(app_id, app_secret).text

    # Set our goramldir client to use JWT token from itsyou.online
    client.api.set_auth_header("Bearer " + jwt_token)

    # try to make simple GET call to goramldir server
    resp = client.api.users.users_get()
    print("resp body =",resp.text)

if __name__ == "__main__":
    '''
    usage : python3 main.py application_id application_secret
    '''
    main(sys.argv[1], sys.argv[2])
