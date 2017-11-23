
from .Address import Address
from .City import City
from .api_response import APIResponse
from .unhandled_api_error import UnhandledAPIError
from .unmarshall_error import UnmarshallError


class UsersService:
    def __init__(self, client):
        self.client = client


    def users_byUserId_address_byAddressId_get(self, addressId, userId, headers=None, query_params=None, content_type="application/json"):
        """
        get address id
        of address
        It is method for GET /users/{userId}/address/{addressId}
        """
        uri = self.client.base_url + "/users/"+userId+"/address/"+addressId
        resp = self.client.get(uri, None, headers, query_params, content_type)
        try:
            if resp.status_code == 200:
                return APIResponse(data=Address(resp.json()), response=resp)

            message = 'unknown status code={}'.format(resp.status_code)
            raise UnhandledAPIError(response=resp, code=resp.status_code,
                                    message=message)
        except ValueError as msg:
            raise UnmarshallError(resp, msg)
        except UnhandledAPIError as uae:
            raise uae
        except Exception as e:
            raise UnmarshallError(resp, e.message)

    def users_byUserId_delete(self, userId, headers=None, query_params=None, content_type="application/json"):
        """
        It is method for DELETE /users/{userId}
        """
        uri = self.client.base_url + "/users/"+userId
        return self.client.delete(uri, None, headers, query_params, content_type)

    def getuserid(self, userId, headers=None, query_params=None, content_type="application/json"):
        """
        get id
        It is method for GET /users/{userId}
        """
        uri = self.client.base_url + "/users/"+userId
        resp = self.client.get(uri, None, headers, query_params, content_type)
        try:
            if resp.status_code == 200:
                return APIResponse(data=City(resp.json()), response=resp)

            message = 'unknown status code={}'.format(resp.status_code)
            raise UnhandledAPIError(response=resp, code=resp.status_code,
                                    message=message)
        except ValueError as msg:
            raise UnmarshallError(resp, msg)
        except UnhandledAPIError as uae:
            raise uae
        except Exception as e:
            raise UnmarshallError(resp, e.message)

    def users_byUserId_post(self, data, userId, headers=None, query_params=None, content_type="application/json"):
        """
        post without request body
        It is method for POST /users/{userId}
        """
        uri = self.client.base_url + "/users/"+userId
        return self.client.post(uri, data, headers, query_params, content_type)

    def users_delete(self, data, headers=None, query_params=None, content_type="application/json"):
        """
        delete with request body
        It is method for DELETE /users
        """
        uri = self.client.base_url + "/users"
        resp = self.client.delete(uri, data, headers, query_params, content_type)
        try:
            if resp.status_code == 200:
                return APIResponse(data=City(resp.json()), response=resp)

            message = 'unknown status code={}'.format(resp.status_code)
            raise UnhandledAPIError(response=resp, code=resp.status_code,
                                    message=message)
        except ValueError as msg:
            raise UnmarshallError(resp, msg)
        except UnhandledAPIError as uae:
            raise uae
        except Exception as e:
            raise UnmarshallError(resp, e.message)

    def get_users(self, data, headers=None, query_params=None, content_type="application/json"):
        """
        First line of comment.
        Second line of comment
        It is method for GET /users
        """
        uri = self.client.base_url + "/users"
        return self.client.get(uri, data, headers, query_params, content_type)

    def option_users(self, headers=None, query_params=None, content_type="application/json"):
        """
        It is method for OPTIONS /users
        """
        uri = self.client.base_url + "/users"
        return self.client.session.options(uri, None, headers, query_params, content_type)

    def create_users(self, data, headers=None, query_params=None, content_type="application/json"):
        """
        create users
        It is method for POST /users
        """
        uri = self.client.base_url + "/users"
        resp = self.client.post(uri, data, headers, query_params, content_type)
        try:
            if resp.status_code == 200:
                return APIResponse(data=City(resp.json()), response=resp)

            message = 'unknown status code={}'.format(resp.status_code)
            raise UnhandledAPIError(response=resp, code=resp.status_code,
                                    message=message)
        except ValueError as msg:
            raise UnmarshallError(resp, msg)
        except UnhandledAPIError as uae:
            raise uae
        except Exception as e:
            raise UnmarshallError(resp, e.message)
