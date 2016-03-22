from flask import Blueprint, jsonify, request
from oauth2_oauth_2_0_queryMwr import *
from oauth2_oauth_2_0_query_ADMINISTRATORMwr import *
from oauth2_oauth_2_0_headerMwr import *


deliveries_api = Blueprint('deliveries_api', __name__)


@deliveries_api.route('/deliveries', methods=['GET'])
@oauth2_oauth_2_0_queryMwr
@oauth2_oauth_2_0_query_ADMINISTRATORMwr
def deliveries_get():
    '''
    Get a list of deliveries
    It is handler for GET /deliveries
    '''
    return jsonify()


@deliveries_api.route('/deliveries', methods=['POST'])
@oauth2_oauth_2_0_headerMwr
def deliveries_post():
    '''
    Create/request a new delivery
    It is handler for POST /deliveries
    '''
    return jsonify()


@deliveries_api.route('/deliveries/<deliveryId>', methods=['GET'])
def deliveries_byDeliveryId_get(deliveryId):
    '''
    Get information on a specific delivery
    It is handler for GET /deliveries/<deliveryId>
    '''
    return jsonify()


@deliveries_api.route('/deliveries/<deliveryId>', methods=['PATCH'])
@oauth2_oauth_2_0_headerMwr
def deliveries_byDeliveryId_patch(deliveryId):
    '''
    Update the information on a specific delivery
    It is handler for PATCH /deliveries/<deliveryId>
    '''
    return jsonify()


@deliveries_api.route('/deliveries/<deliveryId>', methods=['DELETE'])
@oauth2_oauth_2_0_headerMwr
def deliveries_byDeliveryId_delete(deliveryId):
    '''
    Cancel a specific delivery
    It is handler for DELETE /deliveries/<deliveryId>
    '''
    return jsonify()
