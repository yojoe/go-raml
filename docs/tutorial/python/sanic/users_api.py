import json as JSON

from sanic.response import json, text
import jsonschema
from jsonschema import Draft4Validator

User_schema =  JSON.load(open('./schema/User_schema.json'))


async def users_get(request):
    '''
    Get list of all developers
    It is handler for GET /users
    '''
    
    return json({})

async def users_post(request):
    '''
    Add user
    It is handler for POST /users
    '''
    
    inputs = request.json
    try:
        Draft4Validator(User_schema).validate(inputs)
    except jsonschema.ValidationError as e:
        return text('Bad Request Body', 400)
    
    return json({})

async def users_byUsername_get(request, username):
    '''
    Get information on a specific user
    It is handler for GET /users/<username>
    '''
    
    return json({})

