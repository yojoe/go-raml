package iyo

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

const (
	// itsyou.online public key
	iyoPubKey = `
-----BEGIN PUBLIC KEY-----
MHYwEAYHKoZIzj0CAQYFK4EEACIDYgAES5X8XrfKdx9gYayFITc89wad4usrk0n2
7MjiGYvqalizeSWTHEpnd7oea9IQ8T5oJjMVH5cc0H5tFSKilFFeh//wngxIyny6
6+Vq5t5B0V0Ehy01+2ceEon2Y0XDkIKv
-----END PUBLIC KEY-----`
)

var JWTPublicKey *ecdsa.PublicKey

func init() {
	var err error
	JWTPublicKey, err = jwt.ParseECPublicKeyFromPEM([]byte(iyoPubKey))
	if err != nil {
		log.Fatalf("failed to parse pub key:%v", err)
	}

}

// get itsyou.online user's scopes
func getIyoUserScope(tokenStr string) ([]string, error) {
	jwtStr := strings.TrimSpace(strings.TrimPrefix(tokenStr, "token"))
	token, err := jwt.Parse(jwtStr, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodES384 {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return JWTPublicKey, nil
	})
	if err != nil {
		return []string{}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		return []string{}, fmt.Errorf("invalid token")
	}

	// check if the issuer is itsyou.online
	if claims["iss"].(string) != "itsyouonline" {
		return []string{}, fmt.Errorf("invalid issuer:%v", claims["iss"])
	}

	var scopes []string
	for _, v := range claims["scope"].([]interface{}) {
		scopes = append(scopes, v.(string))
	}
	return scopes, nil
}
