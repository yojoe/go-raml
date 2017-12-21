import json as JSON
import jsonschema
from jsonschema import Draft4Validator
from flask import Blueprint, jsonify, request


import os
dir_path = os.path.dirname(os.path.realpath(__file__))

List_Animal_schema = JSON.load(open(dir_path + '/schema/List_Animal_schema.json'))
List_Animal_schema_resolver = jsonschema.RefResolver('file://' + dir_path + '/schema/', List_Animal_schema)
List_Animal_schema_validator = Draft4Validator(List_Animal_schema, resolver=List_Animal_schema_resolver)


animals_api = Blueprint('animals_api', __name__)


@animals_api.route('/animals', methods=['POST'])
def animals_post():
    '''
    create a user
    It is handler for POST /animals
    '''

    inputs = request.get_json()
    try:
        List_Animal_schema_validator.validate(inputs)
    except jsonschema.ValidationError as e:
        return jsonify(errors="bad request body"), 400

    return jsonify()
