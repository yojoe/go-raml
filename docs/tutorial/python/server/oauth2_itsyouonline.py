from functools import wraps
from flask import g, request, jsonify

import jwt

iyo_pub_key = """
-----BEGIN PUBLIC KEY-----
MHYwEAYHKoZIzj0CAQYFK4EEACIDYgAES5X8XrfKdx9gYayFITc89wad4usrk0n2
7MjiGYvqalizeSWTHEpnd7oea9IQ8T5oJjMVH5cc0H5tFSKilFFeh//wngxIyny6
6+Vq5t5B0V0Ehy01+2ceEon2Y0XDkIKv
-----END PUBLIC KEY-----"""

token_prefix = "token "

def get_jwt_scopes(token):
    if token.startswith(token_prefix):
        token = token[len(token_prefix):]
        return jwt.decode(token, iyo_pub_key, verify=False)["scope"]
    else:
        raise Exception('invalid token')


class oauth2_itsyouonline:
    def __init__(self, scopes=None):
        
        self.described_by = "headers"
        self.field = "Authorization"
        
        self.allowed_scopes = scopes

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

            # provide code to check scopes of the access_token
            scopes = get_jwt_scopes(token)

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
