import requests

BASE_URI = "http://localhost:5000"


class Client:
    def __init__(self):
        self.url = BASE_URI
        self.session = requests.Session()
        self.auth_header = ''
    
    def set_auth_header(self, val):
        ''' set authorization header value'''
        self.auth_header = val


    def users_get(self, headers=None, query_params=None):
        """
        Get list of all developers
        It is method for GET /users
        """
        if self.auth_header:
            self.session.headers.update({"Authorization":self.auth_header})

        uri = self.url + "/users"
        return self.session.get(uri, headers=headers, params=query_params)


    def users_post(self, data, headers=None, query_params=None):
        """
        Add user
        It is method for POST /users
        """
        if self.auth_header:
            self.session.headers.update({"Authorization":self.auth_header})

        uri = self.url + "/users"
        return self.session.post(uri, data, headers=headers, params=query_params)


    def users_byUsername_get(self, username, headers=None, query_params=None):
        """
        Get information on a specific user
        It is method for GET /users/{username}
        """
        if self.auth_header:
            self.session.headers.update({"Authorization":self.auth_header})

        uri = self.url + "/users/"+username
        return self.session.get(uri, headers=headers, params=query_params)
