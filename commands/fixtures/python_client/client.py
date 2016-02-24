import requests
from client_utils import build_query_string

BASE_URI = "http://api.jumpscale.com/v3"


class Client:
    def __init__(self):
        self.url = BASE_URI

    def users_get(self, headers=None, query_params=None):
        """
        Get a list of test, this comment is very long, to test our comment generator. Is it
        good?
        It is method for GET /users
        """
        uri = self.url + "/users"
        uri = uri + build_query_string(query_params)
        return requests.get(uri, headers=headers)

    def users_post(self, data, headers=None, query_params=None):
        """
        create users
        It is method for POST /users
        """
        uri = self.url + "/users"
        uri = uri + build_query_string(query_params)
        return requests.post(uri, data, headers=headers)

    def users_byUserId_get(self, userId, headers=None, query_params=None):
        """
        get id
        It is method for GET /users/{userId}
        """
        uri = self.url + "/users/"+userId
        uri = uri + build_query_string(query_params)
        return requests.get(uri, headers=headers)

    def users_byUserId_delete(self, userId, headers=None, query_params=None):
        """
        It is method for DELETE /users/{userId}
        """
        uri = self.url + "/users/"+userId
        uri = uri + build_query_string(query_params)
        return requests.delete(uri, headers=headers)

    def users_byUserId_address_byAddressId_get(self, addressId, userId, headers=None, query_params=None):
        """
        get address id
        It is method for GET /users/{userId}/address/{addressId}
        """
        uri = self.url + "/users/"+userId+"/address/"+addressId
        uri = uri + build_query_string(query_params)
        return requests.get(uri, headers=headers)
