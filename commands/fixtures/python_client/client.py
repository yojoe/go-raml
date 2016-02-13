
import requests

BASE_URI = "http://api.jumpscale.com/v3"

class Client:
    def __init__(self,url):
        self.url = BASE_URI



	def UsersGet(self,headers=None,queryParams=None):
		uri = "/users"
		return requests.get(self.url + uri,headers=headers)


	def UsersPost(self,data,headers=None,queryParams=None):
		uri = "/users"
		return requests.post(self.url + uri, data,headers=headers)


	def UsersUserIdGet(self,userId,headers=None,queryParams=None):
		uri = "/users/"+userId
		return requests.get(self.url + uri,headers=headers)


	def UsersUserIdDelete(self,userId,headers=None,queryParams=None):
		uri = "/users/"+userId
		return requests.delete(self.url + uri,headers=headers)


	def UsersUserIdAddressAddressIdGet(self,addressId,userId,headers=None,queryParams=None):
		uri = "/users/"+userId+"/address/"+addressId
		return requests.get(self.url + uri,headers=headers)


