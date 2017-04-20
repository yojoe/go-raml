class UsersService:
    def __init__(self, client):
        self.client = client



    async def get_users(self, headers=None, query_params=None, content_type="application/json"):
        """
        First line of comment.
        Second line of comment
        It is method for GET /users
        """
        uri = self.client.base_url + "/users"
        return await self.client.get(uri, headers, query_params, content_type)


    async def create_users(self, data, headers=None, query_params=None, content_type="application/json"):
        """
        create users
        It is method for POST /users
        """
        uri = self.client.base_url + "/users"
        return await self.client.post(uri, data, headers, query_params, content_type)


    async def option_users(self, headers=None, query_params=None, content_type="application/json"):
        """
        It is method for OPTIONS /users
        """
        uri = self.client.base_url + "/users"
        return await self.client.options(uri, headers, query_params, content_type)


    async def getuserid(self, userId, headers=None, query_params=None, content_type="application/json"):
        """
        get id
        It is method for GET /users/{userId}
        """
        uri = self.client.base_url + "/users/"+userId
        return await self.client.get(uri, headers, query_params, content_type)


    async def users_byUserId_delete(self, userId, headers=None, query_params=None, content_type="application/json"):
        """
        It is method for DELETE /users/{userId}
        """
        uri = self.client.base_url + "/users/"+userId
        return await self.client.delete(uri, headers, query_params, content_type)


    async def users_byUserId_address_byAddressId_get(self, addressId, userId, headers=None, query_params=None, content_type="application/json"):
        """
        get address id
        of address
        It is method for GET /users/{userId}/address/{addressId}
        """
        uri = self.client.base_url + "/users/"+userId+"/address/"+addressId
        return await self.client.get(uri, headers, query_params, content_type)
