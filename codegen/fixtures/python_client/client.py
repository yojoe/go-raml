import requests
from client_utils import build_query_string

BASE_URI = "http://api.jumpscale.com/v3"


class Client:
    def __init__(self):
        self.url = BASE_URI
        self.session = requests.Session()
        self.auth_header = ''
    
    def set_auth_header(val):
        ''' set authorization header value'''
        self.auth_header = val


    def get_users(self, headers=None, query_params=None):
        """
        First line of comment.
        Second line of comment
        It is method for GET /users
        """
        if self.auth_header:
            if not headers:
                headers = {'Authorization': self.auth_header}
            else:
                headers['Authorization'] = self.auth_header

        uri = self.url + "/users"
        return self.session.get(uri, headers=headers, params=query_params)


    def create_users(self, data, headers=None, query_params=None):
        """
        create users
        It is method for POST /users
        """
        if self.auth_header:
            if not headers:
                headers = {'Authorization': self.auth_header}
            else:
                headers['Authorization'] = self.auth_header

        uri = self.url + "/users"
        return self.session.post(uri, data, headers=headers, params=query_params)


    def getuserid(self, userId, headers=None, query_params=None):
        """
        get id
        It is method for GET /users/{userId}
        """
        if self.auth_header:
            if not headers:
                headers = {'Authorization': self.auth_header}
            else:
                headers['Authorization'] = self.auth_header

        uri = self.url + "/users/"+userId
        return self.session.get(uri, headers=headers, params=query_params)


    def users_byUserId_delete(self, userId, headers=None, query_params=None):
        """
        It is method for DELETE /users/{userId}
        """
        if self.auth_header:
            if not headers:
                headers = {'Authorization': self.auth_header}
            else:
                headers['Authorization'] = self.auth_header

        uri = self.url + "/users/"+userId
        return self.session.delete(uri, headers=headers, params=query_params)


    def users_byUserId_address_byAddressId_get(self, addressId, userId, headers=None, query_params=None):
        """
        get address id
        of address
        It is method for GET /users/{userId}/address/{addressId}
        """
        if self.auth_header:
            if not headers:
                headers = {'Authorization': self.auth_header}
            else:
                headers['Authorization'] = self.auth_header

        uri = self.url + "/users/"+userId+"/address/"+addressId
        return self.session.get(uri, headers=headers, params=query_params)
