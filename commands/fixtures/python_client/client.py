
import requests

BASE_URI = "http://api.jumpscale.com/v3"

class Client:
    def __init__(self,url):
        self.url = BASE_URI



	def UsersGet(self):
		uri = "/users"
		return requests.get(self.url + uri)


	def UsersPost(self,data):
		uri = "/users"
		return requests.post(self.url + uri, data)


	def UsersUserIdGet(self,userId):
		uri = "/users/"+userId
		return requests.get(self.url + uri)


	def UsersUserIdDelete(self,userId):
		uri = "/users/"+userId
		return requests.delete(self.url + uri)


	def UsersUserIdAddressAddressIdGet(self,addressId,userId):
		uri = "/users/"+userId+"/address/"+addressId
		return requests.get(self.url + uri)


