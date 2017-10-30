import json as JSON
import jsonschema
from jsonschema import Draft4Validator
from flask import Blueprint, jsonify, request


import os
dir_path = os.path.dirname(os.path.realpath(__file__))


deliveries_api = Blueprint('deliveries_api', __name__)


@deliveries_api.route('/deliveries', methods=['GET'])
def deliveries_get():
    '''
    Get a list of deliveries
    It is handler for GET /deliveries
    '''
    
    return jsonify()


@deliveries_api.route('/deliveries', methods=['POST'])
def deliveries_post():
    '''
    Create/request a new delivery
    It is handler for POST /deliveries
    '''
    
    return jsonify()


@deliveries_api.route('/deliveries/<deliveryId>', methods=['GET'])
def getDeliveriesByDeliveryID(deliveryId):
    '''
    Get information on a specific delivery
    It is handler for GET /deliveries/<deliveryId>
    '''
    
    return jsonify()


@deliveries_api.route('/deliveries/<deliveryId>', methods=['PATCH'])
def deliveries_byDeliveryId_patch(deliveryId):
    '''
    Update the information on a specific delivery
    It is handler for PATCH /deliveries/<deliveryId>
    '''
    
    return jsonify()


@deliveries_api.route('/deliveries/<deliveryId>', methods=['DELETE'])
def deliveries_byDeliveryId_delete(deliveryId):
    '''
    Cancel a specific delivery
    It is handler for DELETE /deliveries/<deliveryId>
    '''
    
    return jsonify()
