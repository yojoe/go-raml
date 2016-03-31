from flask import Blueprint, jsonify, request

#from UsersPostReqBody import UsersPostReqBody

from UsersPostReqBody import UsersPostReqBody


users_api = Blueprint('users_api', __name__)


@users_api.route('/users', methods=['POST'])
def users_post():
    '''
    create a user
    It is handler for POST /users
    '''
    
    inputs = UsersPostReqBody.from_json(request.get_json())
    if not inputs.validate():
        return jsonify(errors=inputs.errors), 400
    
    return jsonify()


@users_api.route('/users/<id>', methods=['GET'])
def users_byId_get(id):
    '''
    get id of users.
    This method will be return single user object.
    Replace ID with user defined field.
    It is handler for GET /users/<id>
    '''
    
    return jsonify()
