from flask import Blueprint, jsonify, request
import libraries.security.oauth2_Dropbox as oauth2_Dropbox



configs_api = Blueprint('configs_api', __name__)


@configs_api.route('/configs', methods=['GET'])
@oauth2_Dropbox.oauth2_Dropbox([])
def configs_get():
    '''
    get config files
    It is handler for GET /configs
    '''
    
    return jsonify()


@configs_api.route('/configs', methods=['POST'])
@oauth2_Dropbox.oauth2_Dropbox([])
def configs_post():
    '''
    It is handler for POST /configs
    '''
    
    return jsonify()


@configs_api.route('/configs', methods=['PUT'])
@oauth2_Dropbox.oauth2_Dropbox([])
def configs_put():
    '''
    It is handler for PUT /configs
    '''
    
    return jsonify()
