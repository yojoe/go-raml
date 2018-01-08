
from .users_service import UsersService
from .oauth2_client_itsyouonline import Oauth2ClientItsyouonline

from .http_client import HTTPClient


class Client:
    def __init__(self, loop, base_uri="http://localhost:5000", token=None):
        http_client = HTTPClient(loop, base_uri, token)
        self.security_schemes = Security()

        self.users = UsersService(http_client)
        self.close = http_client.close


class Security:
    def __init__(self):
        self.oauth2_client_itsyouonline = Oauth2ClientItsyouonline()
