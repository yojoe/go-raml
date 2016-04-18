import requests
from client_utils import build_query_string

BASE_URI = ""


class Client:
    def __init__(self):
        self.url = BASE_URI

    def resources_get(self, headers=None, query_params=None):
        """
        Get a resource
        It is method for GET /resources
        """
        uri = self.url + "/resources"
        uri = uri + build_query_string(query_params)
        return requests.get(uri, headers=headers)

    def resources_post(self, data, headers=None, query_params=None):
        """
        Post a resource
        It is method for POST /resources
        """
        uri = self.url + "/resources"
        uri = uri + build_query_string(query_params)
        return requests.post(uri, data, headers=headers)
