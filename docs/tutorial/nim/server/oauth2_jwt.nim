import marshal, strutils

import libjwt, jester

type
  Oauth2JWT* = object
    pubKey*: string

  JWTGrants* = object
    iss: string
    globalid: string
    aud: seq[string]
    username: string
    exp: int
    scope: seq[string]

proc decodeJWT*(ojwt: Oauth2JWT, token: string): JWTGrants =
  # decode JWT token and extracts it's grants
  var j: ptr jwt_t
  
  let ret = jwt_decode(addr j, token, ojwt.pubKey, cint(ojwt.pubKey.len))
  if ret > 0:
    return
  result = to[JWTGrants]($(json_dumps(j.grants,0)))

proc checkScopes(s1: openArray[string], s2: openArray[string]):bool =
  #check if at least one element of 1 is member of s2
  if s2.len == 0:
    return true
  for v1 in s1:
    for v2 in s2:
      if v1 == v2:
        return true
  return false


proc checkJWTToken*(ojwt: Oauth2JWT, req: Request, scopes: openArray[string]): bool =
  # check if a JWT token has the supplied scopes
  let authHdr = req.headers.getOrDefault("Authorization")
  
  if authHdr.len == 0:
    return false
  
  if not authHdr.startsWith("token "):
    return false

  var grants: JWTGrants
  grants = ojwt.decodeJWT(authHdr[6..authHdr.len])
  if grants.scope.len == 0: grants.scope = @[]
  
  return checkScopes(grants.scope, scopes)