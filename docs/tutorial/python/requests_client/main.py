import sys
import json

import goramldir
import goramldir.User as User

# goramldir client
client = goramldir.Client()


def main(app_id, app_secret):
    # get JWT token from itsyouonline
    jwt_token = client.oauth2_client_itsyouonline.get_access_token(app_id, app_secret).text

    # Set our goramldir client to use JWT token from itsyou.online
    client.oauth2_client_itsyouonline.set_auth_header("Bearer " + jwt_token)

    # try to make simple GET call to goramldir server
    resp = client.users.users_get()
    print("resp body =",resp.text)

    # example on how to create object and encode it to json
    user = User(name="iwan", username="ibk")
    json_str = user.as_json()
    print(json_str)

    # example on how to create object from json object
    user2 = User(json = json.loads(json_str))
    print(user2.as_json())


if __name__ == "__main__":
    '''
    usage : python3 main.py application_id application_secret
    '''
    main(sys.argv[1], sys.argv[2])
