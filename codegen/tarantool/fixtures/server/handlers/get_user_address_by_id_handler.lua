-- THIS FILE IS SAFE TO EDIT. It will not be overwritten when rerunning go-raml.

local schema = require("handlers.schemas.schema")

function get_user_address_by_id_handler(request)
    -- handler for GET /users/:userId/address/:addressId
    -- response body for 200 should match schema.Address

    local resp = {
    }
    return resp
end

return get_user_address_by_id_handler