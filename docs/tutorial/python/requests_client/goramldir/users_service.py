class UsersService:
    def __init__(self, client):
        self.client = client



    def users_get(self, headers=None, query_params=None):
        """
        Get list of all developers
        It is method for GET /users
        """
        uri = self.client.base_url + "/users"
        return self.client.session.get(uri, headers=headers, params=query_params)


    def users_post(self, data, headers=None, query_params=None):
        """
        Add user
        It is method for POST /users
        """
        uri = self.client.base_url + "/users"
        return self.client.post(uri, data, headers=headers, params=query_params)


    def users_byUsername_get(self, username, headers=None, query_params=None):
        """
        Get information on a specific user
        It is method for GET /users/{username}
        """
        uri = self.client.base_url + "/users/"+username
        return self.client.session.get(uri, headers=headers, params=query_params)
