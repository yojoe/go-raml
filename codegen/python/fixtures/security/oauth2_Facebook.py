from functools import wraps
from flask import g, request, jsonify

from jose import jwt

oauth2_server_pub_key = """"""

token_prefix = "Bearer "

def get_jwt_scopes(token, audience):
    if token.startswith(token_prefix):
        token = token[len(token_prefix):]
        return jwt.decode(token, oauth2_server_pub_key, audience=audience)["scope"]
    else:
        raise Exception('invalid token')

class oauth2_Facebook:
    def __init__(self, scopes=None, audience= None):
        
        self.described_by = "headers"
        self.field = "Authorization"
        
        self.allowed_scopes = scopes
        if audience is None:
            self.audience = ''
        else:
            self.audience = ",".join(audience)

    def __call__(self, f):
        @wraps(f)
        def decorated_function(*args, **kwargs):
            token = ""
            if self.described_by == "headers":
                token = request.headers.get(self.field, "")
            elif self.described_by == "queryParameters":
                token = request.args.get("access_token", "")

            if token == "":
                return jsonify(), 401

            g.access_token = token

            if len(oauth2_server_pub_key) > 0:
                scopes = get_jwt_scopes(token, self.audience)
                if self.check_scopes(scopes) == False:
                    return jsonify(), 403
            return f(*args, **kwargs)
        return decorated_function

    def check_scopes(self, scopes):
        if self.allowed_scopes is None or len(self.allowed_scopes) == 0:
            return True

        for allowed in self.allowed_scopes:
            for s in scopes:
                if s == allowed:
                    return True

        return False