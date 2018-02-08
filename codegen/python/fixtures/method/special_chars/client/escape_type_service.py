# DO NOT EDIT THIS FILE. This file will be overwritten when re-running go-raml.
from .User2_0 import User2_0
from .unhandled_api_error import UnhandledAPIError
from .unmarshall_error import UnmarshallError


class Escape_typeService:
    def __init__(self, client):
        self.client = client

    def escape_type_post(self, data, headers=None, query_params=None, content_type="application/json"):
        """
        It is method for POST /escape_type
        """
        if query_params is None:
            query_params = {}

        uri = self.client.base_url + "/escape_type"
        resp = self.client.post(uri, data, headers, query_params, content_type)
        try:
            if resp.status_code == 200:
                return User2_0(resp.json()), resp

            message = 'unknown status code={}'.format(resp.status_code)
            raise UnhandledAPIError(response=resp, code=resp.status_code,
                                    message=message)
        except ValueError as msg:
            raise UnmarshallError(resp, msg)
        except UnhandledAPIError as uae:
            raise uae
        except Exception as e:
            raise UnmarshallError(resp, e.message)
