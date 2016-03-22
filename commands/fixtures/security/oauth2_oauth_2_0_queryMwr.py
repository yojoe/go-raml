from functools import wraps
from flask import g, request, jsonify

def oauth2_oauth_2_0_queryMwr(f):
    @wraps(f)
    def decorated_function(*args, **kwargs):
        token = request.args.get("access_token", "")
        
        if token == "":
            return jsonify(), 401

        g.access_token = token

        return f(*args, **kwargs)
    return decorated_function
