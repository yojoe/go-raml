-- THIS FILE IS SAFE TO EDIT. It will not be overwritten when rerunning go-raml.

local schema = require("handlers.schemas.schema")

function usersusername_get_handler(request)
    -- handler for GET /users/:username
    -- response body for 200 should match schema.User

    local resp = {
    }
    return resp
end

return usersusername_get_handler