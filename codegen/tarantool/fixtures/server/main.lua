
require("handlers.helloworld_handler")
require("handlers.users_handler")
require("handlers.usersuserId_handler")
require("handlers.usersuserIdaddressaddressId_handler")

httpd = require('http.server').new("0.0.0.0", 5000)


httpd:route({ path = '/helloworld', method = 'GET'}, helloworld_get)

httpd:route({ path = '/users', method = 'GET'}, users_get)
httpd:route({ path = '/users', method = 'POST'}, users_post)

httpd:route({ path = '/users/:userId', method = 'GET'}, usersuser_id_get)
httpd:route({ path = '/users/:userId', method = 'DELETE'}, usersuser_id_delete)

httpd:route({ path = '/users/:userId/address/:addressId', method = 'GET'}, get_user_address_by_id)

httpd:start()