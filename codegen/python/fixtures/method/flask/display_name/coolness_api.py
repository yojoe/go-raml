import json as JSON
import jsonschema
from jsonschema import Draft4Validator
from flask import Blueprint, jsonify, request


import os
dir_path = os.path.dirname(os.path.realpath(__file__))

CreateSomethingCoolReqBody_schema = JSON.load(open(dir_path + '/schema/CreateSomethingCoolReqBody_schema.json'))
CreateSomethingCoolReqBody_schema_resolver = jsonschema.RefResolver(
    'file://' + dir_path + '/schema/', CreateSomethingCoolReqBody_schema)
CreateSomethingCoolReqBody_schema_validator = Draft4Validator(
    CreateSomethingCoolReqBody_schema,
    resolver=CreateSomethingCoolReqBody_schema_resolver)


coolness_api = Blueprint('coolness_api', __name__)


@coolness_api.route('/coolness', methods=['POST'])
def CreateSomethingCool():
    '''
    It is handler for POST /coolness
    '''

    inputs = request.get_json()
    try:
        CreateSomethingCoolReqBody_schema_validator.validate(inputs)
    except jsonschema.ValidationError as e:
        return jsonify(errors="bad request body"), 400

    return jsonify()
