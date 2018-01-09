import sys
import asyncio

import goramldir



async def main(app_id, app_secret, loop):
    client = goramldir.Client(loop)
    # get JWT token from itsyouonline
    resp = await client.oauth2_client_itsyouonline.get_access_token(app_id, app_secret)
    jwt_token = await resp.text()
    print("jwt_token=", jwt_token)

    # Set our goramldir client to use JWT token from itsyou.online
    client.oauth2_client_itsyouonline.set_auth_header("Bearer " + jwt_token)

    # try to make simple GET call to goramldir server
    resp = await client.users.users_get()
    print("resp body =",await resp.text())

    client.close()
if __name__ == "__main__":
    '''
    usage : python3 main.py application_id application_secret
    '''
    loop = asyncio.get_event_loop()
    loop.run_until_complete(main(sys.argv[1], sys.argv[2], loop))
