from flask import Blueprint, jsonify, request


from CreateSomethingCoolReqBody import CreateSomethingCoolReqBody

coolness_api = Blueprint('coolness_api', __name__)


@coolness_api.route('/coolness', methods=['POST'])
def CreateSomethingCool():
    '''
    It is handler for POST /coolness
    '''
    
    inputs = CreateSomethingCoolReqBody.from_json(request.get_json())
    if not inputs.validate():
        return jsonify(errors=inputs.errors), 400
    
    return jsonify()
