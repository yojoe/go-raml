-- THIS FILE IS SAFE TO EDIT. It will not be overwritten when rerunning go-raml.

local schema = require("handlers.schemas.schema")

function users_post_handler(request)
    -- handler for POST /users
    -- request body should match schema.User
    -- response body for 200 should match schema.User

    local resp = {
    }
    return resp
end

return users_post_handler