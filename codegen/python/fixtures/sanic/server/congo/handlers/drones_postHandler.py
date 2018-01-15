# THIS FILE IS SAFE TO EDIT. It will not be overwritten when rerunning go-raml.

from sanic.response import json, text

import json as JSON
import jsonschema
from jsonschema import Draft4Validator

import os

dir_path = os.path.dirname(os.path.realpath(__file__))
User_schema = JSON.load(open(dir_path + '/schema/User_schema.json'))
User_schema_resolver = jsonschema.RefResolver('file://' + dir_path + '/schema/', User_schema)
User_schema_validator = Draft4Validator(User_schema, resolver=User_schema_resolver)


def drones_postHandler(request):

    inputs = request.json

    try:
        User_schema_validator.validate(inputs)
    except jsonschema.ValidationError as e:
        return text('Bad Request Body', 400)

    return json({})
