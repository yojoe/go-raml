import requests
from client_utils import build_query_string

BASE_URI = ""


class Client:
    def __init__(self):
        self.url = BASE_URI

    def organisation_get(self, headers=None, query_params=None):
        """
        Returns an organisation entity.
        It is method for GET /organisation
        """
        uri = self.url + "/organisation"
        uri = uri + build_query_string(query_params)
        return requests.get(uri, headers=headers)

    def organisation_post(self, data, headers=None, query_params=None):
        """
        It is method for POST /organisation
        """
        uri = self.url + "/organisation"
        uri = uri + build_query_string(query_params)
        return requests.post(uri, data, headers=headers)
