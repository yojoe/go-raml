-- THIS FILE IS SAFE TO EDIT. It will not be overwritten when rerunning go-raml.

local schema = require("handlers.schemas.schema")

function users_get_handler(request)
    -- handler for GET /users
    -- response body for 200 should match schema.UsersGetRespBody

    local resp = {
    }
    return resp
end

return users_get_handler