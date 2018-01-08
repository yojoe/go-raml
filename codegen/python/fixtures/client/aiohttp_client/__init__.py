
from .users_service import UsersService

from .http_client import HTTPClient


class Client:
    def __init__(self, loop, base_uri="http://api.jumpscale.com/v3", token=None):
        http_client = HTTPClient(loop, base_uri, token)

        self.users = UsersService(http_client)
        self.close = http_client.close
