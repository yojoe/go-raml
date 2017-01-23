import json as JSON

from sanic.response import json, text
import jsonschema
from jsonschema import Draft4Validator

User_schema =  JSON.load(open('./schema/User_schema.json'))


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
        Draft4Validator(User_schema).validate(inputs)
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
        Draft4Validator(User_schema).validate(inputs)
    except jsonschema.ValidationError as e:
        return text('Bad Request Body', 400)
    
    return json({})

async def deliveries_byDeliveryId_delete(request, deliveryId):
    '''
    Cancel a specific delivery
    It is handler for DELETE /deliveries/<deliveryId>
    '''
    
    return json({})

