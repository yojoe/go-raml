from functools import wraps
from flask import g, request, jsonify

def Facebook_ADMINISTRATOR(f):
    @wraps(f)
    def decorated_function(*args, **kwargs):
        allowed_scopes = ["ADMINISTRATOR"]
        scopes = []

        token = g.access_token
	
        # provide code to check scopes of the access_token

        # check scopes
        for allowed in allowed_scopes:
            for s in scopes:
                if s == allowed:
                    return f(*args, **kwargs)
        # scopes doesn't match
        return jsonify(), 401

    return decorated_function
