from functools import wraps
from flask import g, request, jsonify

def oauth2_oauth_2_0_headerMwr(f):
    @wraps(f)
    def decorated_function(*args, **kwargs):
        token = request.headers.get("Authorization", "")
        
        if token == "":
            return jsonify(), 401

        g.access_token = token

        return f(*args, **kwargs)
    return decorated_function
