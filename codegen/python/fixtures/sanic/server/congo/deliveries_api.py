import json as JSON

from sanic.response import json, text
import jsonschema
from jsonschema import Draft4Validator

import os
dir_path = os.path.dirname(os.path.realpath(__file__))

User_schema =  JSON.load(open(dir_path + '/schema/User_schema.json'))
User_schema_resolver = jsonschema.RefResolver('file://' + dir_path + '/schema/', User_schema)
User_schema_validator = Draft4Validator(User_schema, resolver=User_schema_resolver)



async def deliveries_get(request):
    '''
    Get a list of deliveries
    It is handler for GET /deliveries
    '''
    
    return json({})

async def deliveries_post(request):
    '''
    Create/request a new delivery
    It is handler for POST /deliveries
    '''
    
    inputs = request.json
    try:
        User_schema_validator.validate(inputs)
    except jsonschema.ValidationError as e:
        return text('Bad Request Body', 400)
    
    return json({})

async def deliveries_byDeliveryId_get(request, deliveryId):
    '''
    Get information on a specific delivery
    It is handler for GET /deliveries/<deliveryId>
    '''
    
    return json({})

async def deliveries_byDeliveryId_patch(request, deliveryId):
    '''
    Update the information on a specific delivery
    It is handler for PATCH /deliveries/<deliveryId>
    '''
    
    inputs = request.json
    try:
        User_schema_validator.validate(inputs)
    except jsonschema.ValidationError as e:
        return text('Bad Request Body', 400)
    
    return json({})

async def deliveries_byDeliveryId_delete(request, deliveryId):
    '''
    Cancel a specific delivery
    It is handler for DELETE /deliveries/<deliveryId>
    '''
    
    return json({})

