class UsersService:
    def __init__(self, client):
        self.client = client



    async def users_get(self, headers=None, query_params=None):
        """
        Get list of all developers
        It is method for GET /users
        """
        uri = self.client.base_url + "/users"
        return await self.client.get(uri, headers=headers, params=query_params)


    async def users_post(self, data, headers=None, query_params=None):
        """
        Add user
        It is method for POST /users
        """
        uri = self.client.base_url + "/users"
        return await self.client.post(uri, data, headers=headers, params=query_params)


    async def users_byUsername_get(self, username, headers=None, query_params=None):
        """
        Get information on a specific user
        It is method for GET /users/{username}
        """
        uri = self.client.base_url + "/users/"+username
        return await self.client.get(uri, headers=headers, params=query_params)
