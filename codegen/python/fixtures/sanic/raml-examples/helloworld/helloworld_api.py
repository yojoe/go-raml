import json as JSON

from sanic.response import json, text
import jsonschema
from jsonschema import Draft4Validator

import os
dir_path = os.path.dirname(os.path.realpath(__file__))



async def helloworld_get(request):
    '''
    It is handler for GET /helloworld
    '''
    
    return json({})

