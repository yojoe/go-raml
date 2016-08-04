package main

import (
	"flag"
	"log"

	"github.com/itsyouonline/identityserver/clients/go/itsyouonline"
)

var (
	appID     = flag.String("app_id", "", "application ID")
	appSecret = flag.String("app_secret", "", "application secret")
)

func main() {
	flag.Parse()
	if *appID == "" || *appSecret == "" {
		log.Fatalf("please specify itsyou.online application ID & API Key")
	}

	// create itsyou.online client
	ioc := itsyouonline.NewItsyouonline()

	// get oauth2 token
	_, err := ioc.LoginWithClientCredentials(*appID, *appSecret)
	if err != nil {
		log.Fatalf("failed to get itsyou.online token:%v\n", err)
	}

	// create itsyou.online JWT token
	jwtToken, err := ioc.CreateJWTToken([]string{"user:memberof:goraml"}, nil)
	if err != nil {
		log.Fatalf("failed to create itsyou.online JWT token:%v", err)
	}

	// create goramldir client
	gr := Newgoramldir()

	// set goramldir authorization header to use JWT token
	gr.AuthHeader = "token " + jwtToken

	// calling GET /users
	users, resp, err := gr.UsersGet(map[string]interface{}{}, map[string]interface{}{})
	if err != nil {
		log.Fatalf("failed to GET /users:%v, resp code = %v", err, resp.StatusCode)
	}

	if resp.StatusCode != 200 {
		log.Fatalf("GET /users failed. http status code = %v", resp.StatusCode)
	}

	log.Printf("users = %v\n", users)

}
