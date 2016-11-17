import jester, marshal, system
import oauth2_jwt


let ojwt = Oauth2JWT(pubKey:readFile("itsyouonline.pub"))


proc deliveriesGet*(req: Request) : tuple[code: HttpCode, content: string] =
  # Get a list of deliveries
  let respBody = ""
  
  
  result = (code: Http200, content: respBody)

proc deliveriesPost*(req: Request) : tuple[code: HttpCode, content: string] =
  # Create/request a new delivery
  let respBody = ""
  
  
  result = (code: Http200, content: respBody)

proc getDeliveriesByDeliveryID*(deliveryId: string, req: Request) : tuple[code: HttpCode, content: string] =
  # Get information on a specific delivery
  let respBody = ""
  
  
  result = (code: Http200, content: respBody)

proc deliveriesByDeliveryIdPatch*(deliveryId: string, req: Request) : tuple[code: HttpCode, content: string] =
  # Update the information on a specific delivery
  let respBody = ""
  
  
  result = (code: Http200, content: respBody)

proc deliveriesByDeliveryIdDelete*(deliveryId: string, req: Request) : tuple[code: HttpCode, content: string] =
  # Cancel a specific delivery
  let respBody = ""
  
  
  result = (code: Http200, content: respBody)

