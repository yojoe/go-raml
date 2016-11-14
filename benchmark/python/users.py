from flask import Blueprint, jsonify, request


from User import User

users_api = Blueprint('users_api', __name__)


@users_api.route('/users', methods=['GET'])
def users_get():
    '''
    Get random user
    It is handler for GET /users
    '''
    
    return jsonify({"name":"John", "username":"Doe"})


@users_api.route('/users', methods=['POST'])
def users_post():
    '''
    Add user
    It is handler for POST /users
    '''
    
    inputs = User.from_json(request.get_json())
    #if not inputs.validate():
    #    return jsonify(errors=inputs.errors), 400
    return inputs.name.data


@users_api.route('/users/<username>', methods=['GET'])
def users_byUsername_get(username):
    '''
    Get information on a specific user
    It is handler for GET /users/<username>
    '''
    
    return jsonify()
