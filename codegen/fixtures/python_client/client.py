import requests

from .users_service import  UsersService 

BASE_URI = "http://api.jumpscale.com/v3"


class Client:
    def __init__(self):
        self.base_url = BASE_URI
        self.session = requests.Session()
        self.session.headers.update({"Content-Type": "application/json"})
        
        self.users = UsersService(self)
    
    def set_auth_header(self, val):
        ''' set authorization header value'''
        self.session.headers.update({"Authorization":val})