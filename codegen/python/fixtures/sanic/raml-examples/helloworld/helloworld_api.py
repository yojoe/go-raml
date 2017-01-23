import json as JSON

from sanic.response import json, text
import jsonschema
from jsonschema import Draft4Validator



async def helloworld_get(request):
    '''
    It is handler for GET /helloworld
    '''
    
    return json({})

