from flask import Blueprint, jsonify, request

import oauth2_itsyouonline as oauth2_itsyouonline

from User import User

users_api = Blueprint('users_api', __name__)


@users_api.route('/users', methods=['GET'])
@oauth2_itsyouonline.oauth2_itsyouonline([])
def users_get():
    '''
    Get list of all developers
    It is handler for GET /users
    '''
    
    return jsonify()


@users_api.route('/users', methods=['POST'])
@oauth2_itsyouonline.oauth2_itsyouonline(["user:memberof:goraml-admin"])
def users_post():
    '''
    Add user
    It is handler for POST /users
    '''
    
    inputs = User.from_json(request.get_json())
    if not inputs.validate():
        return jsonify(errors=inputs.errors), 400
    
    return jsonify()


@users_api.route('/users/<username>', methods=['GET'])
@oauth2_itsyouonline.oauth2_itsyouonline(["user:memberof:goraml"])
def users_byUsername_get(username):
    '''
    Get information on a specific user
    It is handler for GET /users/<username>
    '''
    
    return jsonify()
