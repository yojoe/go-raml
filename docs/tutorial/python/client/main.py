import itsyouonline

app_id = '...'
api_key = '....'

client = itsyouonline.Client()

resp = client.oauth.LoginViaClientCredentials(client_id=app_id, client_secret=api_key)
print(resp)
