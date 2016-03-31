from flask import Blueprint, jsonify, request


from User import User

from User import User


deliveries_api = Blueprint('deliveries_api', __name__)


@deliveries_api.route('/deliveries', methods=['GET'])
def deliveries_get():
    '''
    Get a list of deliveries
    It is handler for GET /deliveries
    '''
    
    return jsonify()


@deliveries_api.route('/deliveries', methods=['POST'])
def deliveries_post():
    '''
    Create/request a new delivery
    It is handler for POST /deliveries
    '''
    
    inputs = User.from_json(request.get_json())
    if not inputs.validate():
        return jsonify(errors=inputs.errors), 400
    
    return jsonify()


@deliveries_api.route('/deliveries/<deliveryId>', methods=['GET'])
def deliveries_byDeliveryId_get(deliveryId):
    '''
    Get information on a specific delivery
    It is handler for GET /deliveries/<deliveryId>
    '''
    
    return jsonify()


@deliveries_api.route('/deliveries/<deliveryId>', methods=['PATCH'])
def deliveries_byDeliveryId_patch(deliveryId):
    '''
    Update the information on a specific delivery
    It is handler for PATCH /deliveries/<deliveryId>
    '''
    
    inputs = User.from_json(request.get_json())
    if not inputs.validate():
        return jsonify(errors=inputs.errors), 400
    
    return jsonify()


@deliveries_api.route('/deliveries/<deliveryId>', methods=['DELETE'])
def deliveries_byDeliveryId_delete(deliveryId):
    '''
    Cancel a specific delivery
    It is handler for DELETE /deliveries/<deliveryId>
    '''
    
    return jsonify()
