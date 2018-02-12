# THIS FILE IS SAFE TO EDIT. It will not be overwritten when rerunning go-raml.

from flask import jsonify, request

import json as JSON
import jsonschema
from jsonschema import Draft4Validator

import os

dir_path = os.path.dirname(os.path.realpath(__file__))
User_schema = JSON.load(open(dir_path + '/schema/User_schema.json'))
User_schema_resolver = jsonschema.RefResolver('file://' + dir_path + '/schema/', User_schema)
User_schema_validator = Draft4Validator(User_schema, resolver=User_schema_resolver)


def users_postHandler():

    inputs = request.get_json()

    try:
        User_schema_validator.validate(inputs)
    except jsonschema.ValidationError as e:
        return jsonify(errors="bad request body"), 400

    return jsonify(inputs)
