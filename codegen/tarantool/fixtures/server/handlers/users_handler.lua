schema = require("schemas.schema")


function users_get(request)
    -- handler for GET /users
    -- response body for 200 should match schema.UsersGetRespBody

    local resp = {
    }
    return resp
end

function users_post(request)
    -- handler for POST /users
    -- request body should match schema.User
    -- response body for 200 should match schema.User

    local resp = {
    }
    return resp
end
