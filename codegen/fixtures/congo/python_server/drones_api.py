import json as JSON
import jsonschema
from jsonschema import Draft4Validator
from flask import Blueprint, jsonify, request


import os
dir_path = os.path.dirname(os.path.realpath(__file__))

User_schema =  JSON.load(open(dir_path + '/schema/User_schema.json'))
User_schema_resolver = jsonschema.RefResolver('file://' + dir_path + '/schema/', User_schema)
User_schema_validator = Draft4Validator(User_schema, resolver=User_schema_resolver)


drones_api = Blueprint('drones_api', __name__)


@drones_api.route('/drones', methods=['GET'])
def drones_get():
    '''
    Get a list of drones
    It is handler for GET /drones
    '''
    
    return jsonify()


@drones_api.route('/drones', methods=['POST'])
def drones_post():
    '''
    Add a new drone to the fleet
    It is handler for POST /drones
    '''
    
    inputs = request.get_json()
    try:
        User_schema_validator.validate(inputs)
    except jsonschema.ValidationError as e:
        return jsonify(errors="bad request body"), 400
    
    return jsonify()


@drones_api.route('/drones/<droneId>', methods=['GET'])
def drones_byDroneId_get(droneId):
    '''
    Get information on a specific drone
    It is handler for GET /drones/<droneId>
    '''
    
    return jsonify()


@drones_api.route('/drones/<droneId>', methods=['PATCH'])
def drones_byDroneId_patch(droneId):
    '''
    Update the information on a specific drone
    It is handler for PATCH /drones/<droneId>
    '''
    
    inputs = request.get_json()
    try:
        User_schema_validator.validate(inputs)
    except jsonschema.ValidationError as e:
        return jsonify(errors="bad request body"), 400
    
    return jsonify()


@drones_api.route('/drones/<droneId>', methods=['DELETE'])
def drones_byDroneId_delete(droneId):
    '''
    Remove a drone from the fleet
    It is handler for DELETE /drones/<droneId>
    '''
    
    return jsonify()


@drones_api.route('/drones/<droneId>/deliveries', methods=['GET'])
def drones_byDroneId_deliveries_get(droneId):
    '''
    The deliveries scheduled for the current drone
    It is handler for GET /drones/<droneId>/deliveries
    '''
    
    return jsonify()
