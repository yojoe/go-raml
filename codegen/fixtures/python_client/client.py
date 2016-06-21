import requests
from client_utils import build_query_string

BASE_URI = "http://api.jumpscale.com/v3"


class Client:
    def __init__(self):
        self.url = BASE_URI
        self.session = requests.Session()

    def get_users(self, headers=None, query_params=None):
        """
        First line of comment.
        Second line of comment
        It is method for GET /users
        """
        uri = self.url + "/users"
        uri = uri + build_query_string(query_params)
        return self.session.get(uri, headers=headers)

    def create_users(self, data, headers=None, query_params=None):
        """
        create users
        It is method for POST /users
        """
        uri = self.url + "/users"
        uri = uri + build_query_string(query_params)
        return self.session.post(uri, data, headers=headers)

    def getuserid(self, userId, headers=None, query_params=None):
        """
        get id
        It is method for GET /users/{userId}
        """
        uri = self.url + "/users/"+userId
        uri = uri + build_query_string(query_params)
        return self.session.get(uri, headers=headers)

    def users_byUserId_delete(self, userId, headers=None, query_params=None):
        """
        It is method for DELETE /users/{userId}
        """
        uri = self.url + "/users/"+userId
        uri = uri + build_query_string(query_params)
        return self.session.delete(uri, headers=headers)

    def users_byUserId_address_byAddressId_get(self, addressId, userId, headers=None, query_params=None):
        """
        get address id
        of address
        It is method for GET /users/{userId}/address/{addressId}
        """
        uri = self.url + "/users/"+userId+"/address/"+addressId
        uri = uri + build_query_string(query_params)
        return self.session.get(uri, headers=headers)
