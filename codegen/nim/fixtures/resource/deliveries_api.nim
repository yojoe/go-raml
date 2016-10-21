import jester, marshal



proc deliveriesGet*(req: PRequest) : tuple[code: HttpCode, content: string] =
  
  let respBody = ""
  result = (code: Http200, content: respBody)

proc deliveriesPost*(req: PRequest) : tuple[code: HttpCode, content: string] =
  
  let respBody = ""
  result = (code: Http200, content: respBody)

proc getDeliveriesByDeliveryID*(deliveryId: string, req: PRequest) : tuple[code: HttpCode, content: string] =
  
  let respBody = ""
  result = (code: Http200, content: respBody)

proc deliveriesByDeliveryIdPatch*(deliveryId: string, req: PRequest) : tuple[code: HttpCode, content: string] =
  
  let respBody = ""
  result = (code: Http200, content: respBody)

proc deliveriesByDeliveryIdDelete*(deliveryId: string, req: PRequest) : tuple[code: HttpCode, content: string] =
  
  let respBody = ""
  result = (code: Http200, content: respBody)

