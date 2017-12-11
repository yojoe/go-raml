
require("handlers.helloworld_handler")
require("handlers.users_handler")
require("handlers.usersuserId_handler")
require("handlers.usersuserIdaddressaddressId_handler")

httpd = require('http.server').new("0.0.0.0", 5000)


httpd:route({ path = '/helloworld', method = 'GET'}, helloworldGET)

httpd:route({ path = '/users', method = 'GET'}, usersGET)
httpd:route({ path = '/users', method = 'POST'}, usersPOST)

httpd:route({ path = '/users/:userId', method = 'GET'}, usersuserIdGET)
httpd:route({ path = '/users/:userId', method = 'DELETE'}, usersuserIdDELETE)

httpd:route({ path = '/users/:userId/address/:addressId', method = 'GET'}, getUserAddressByID)

httpd:start()