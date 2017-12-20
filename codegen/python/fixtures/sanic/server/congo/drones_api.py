import json as JSON

from sanic.response import json, text
import jsonschema
from jsonschema import Draft4Validator

import os
dir_path = os.path.dirname(os.path.realpath(__file__))

User_schema = JSON.load(open(dir_path + '/schema/User_schema.json'))
User_schema_resolver = jsonschema.RefResolver('file://' + dir_path + '/schema/', User_schema)
User_schema_validator = Draft4Validator(User_schema, resolver=User_schema_resolver)


async def drones_get(request):
    '''
    Get a list of drones
    It is handler for GET /drones
    '''

    return json({})


async def drones_post(request):
    '''
    Add a new drone to the fleet
    It is handler for POST /drones
    '''

    inputs = request.json
    try:
        User_schema_validator.validate(inputs)
    except jsonschema.ValidationError as e:
        return text('Bad Request Body', 400)

    return json({})


async def drones_byDroneId_get(request, droneId):
    '''
    Get information on a specific drone
    It is handler for GET /drones/<droneId>
    '''

    return json({})


async def drones_byDroneId_patch(request, droneId):
    '''
    Update the information on a specific drone
    It is handler for PATCH /drones/<droneId>
    '''

    inputs = request.json
    try:
        User_schema_validator.validate(inputs)
    except jsonschema.ValidationError as e:
        return text('Bad Request Body', 400)

    return json({})


async def drones_byDroneId_delete(request, droneId):
    '''
    Remove a drone from the fleet
    It is handler for DELETE /drones/<droneId>
    '''

    return json({})


async def drones_byDroneId_deliveries_get(request, droneId):
    '''
    The deliveries scheduled for the current drone
    It is handler for GET /drones/<droneId>/deliveries
    '''

    return json({})
