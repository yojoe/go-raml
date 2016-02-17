
import requests
from client_utils import buildQueryString

BASE_URI = "http://api.jumpscale.com/v3"

class Client:
    def __init__(self):
        self.url = BASE_URI



    def UsersGet(self,headers=None,queryParams=None):
        
        """
        Get a list of test 
        """
        
        uri = "/users"
        return requests.get(self.url + uri + buildQueryString(queryParams) ,headers=headers)


    def UsersPost(self,data,headers=None,queryParams=None):
        
        """
        create users 
        """
        
        uri = "/users"
        return requests.post(self.url + uri + buildQueryString(queryParams), data ,headers=headers)


    def UsersUserIdGet(self,userId,headers=None,queryParams=None):
        
        """
        get id 
        """
        
        uri = "/users/"+userId
        return requests.get(self.url + uri + buildQueryString(queryParams) ,headers=headers)


    def UsersUserIdDelete(self,userId,headers=None,queryParams=None):
        
        
        uri = "/users/"+userId
        return requests.delete(self.url + uri + buildQueryString(queryParams) ,headers=headers)


    def UsersUserIdAddressAddressIdGet(self,addressId,userId,headers=None,queryParams=None):
        
        """
        get address id 
        """
        
        uri = "/users/"+userId+"/address/"+addressId
        return requests.get(self.url + uri + buildQueryString(queryParams) ,headers=headers)


