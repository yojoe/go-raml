# THIS FILE IS SAFE TO EDIT. It will not be overwritten when rerunning go-raml.

from flask import jsonify, request

import json as JSON
import jsonschema
from jsonschema import Draft4Validator

import os

dir_path = os.path.dirname(os.path.realpath(__file__))
User2_0_schema = JSON.load(open(dir_path + '/schema/User2_0_schema.json'))
User2_0_schema_resolver = jsonschema.RefResolver('file://' + dir_path + '/schema/', User2_0_schema)
User2_0_schema_validator = Draft4Validator(User2_0_schema, resolver=User2_0_schema_resolver)


def escape_type_postHandler():

    inputs = request.get_json()

    try:
        User2_0_schema_validator.validate(inputs)
    except jsonschema.ValidationError as e:
        return jsonify(errors="bad request body"), 400

    return jsonify()
