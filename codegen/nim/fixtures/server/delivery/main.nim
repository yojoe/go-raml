import jester, asyncdispatch, json, marshal, system

import deliveries_api

routes:
  GET "/deliveries":
    let ret = deliveriesGet(request)
    resp(ret.code, $$ret.content)

  POST "/deliveries":
    let ret = deliveriesPost(request)
    resp(ret.code, $$ret.content)

  GET "/deliveries/@deliveryId":
    let ret = getDeliveriesByDeliveryID(@"deliveryId", request)
    resp(ret.code, $$ret.content)

  PATCH "/deliveries/@deliveryId":
    let ret = deliveriesByDeliveryIdPatch(@"deliveryId", request)
    resp(ret.code, $$ret.content)

  DELETE "/deliveries/@deliveryId":
    let ret = deliveriesByDeliveryIdDelete(@"deliveryId", request)
    resp(ret.code, $$ret.content)


  GET "/":
    resp(readFile("index.html"))

runForever()
