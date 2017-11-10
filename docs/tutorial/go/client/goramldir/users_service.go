package goramldir

import (
	"encoding/json"
	"net/http"

	"github.com/Jumpscale/go-raml/docs/tutorial/go/client/goramldir/goraml"
)

type UsersService service

// Get information on a specific user
func (s *UsersService) UsersUsernameGet(username string, headers, queryParams map[string]interface{}) (User, *http.Response, error) {
	var err error
	var respBody200 User

	resp, err := s.client.doReqNoBody("GET", s.client.BaseURI+"/users/"+username, headers, queryParams)
	if err != nil {
		return respBody200, nil, err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 200:
		err = json.NewDecoder(resp.Body).Decode(&respBody200)
	default:
		err = goraml.NewAPIError(resp, nil)
	}

	return respBody200, resp, err
}

// Get list of all developers
func (s *UsersService) UsersGet(headers, queryParams map[string]interface{}) ([]User, *http.Response, error) {
	var err error
	var respBody200 []User

	resp, err := s.client.doReqNoBody("GET", s.client.BaseURI+"/users", headers, queryParams)
	if err != nil {
		return respBody200, nil, err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 200:
		err = json.NewDecoder(resp.Body).Decode(&respBody200)
	default:
		err = goraml.NewAPIError(resp, nil)
	}

	return respBody200, resp, err
}

// Add user
func (s *UsersService) UsersPost(body User, headers, queryParams map[string]interface{}) (User, *http.Response, error) {
	var err error
	var respBody201 User

	resp, err := s.client.doReqWithBody("POST", s.client.BaseURI+"/users", &body, headers, queryParams)
	if err != nil {
		return respBody201, nil, err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 201:
		err = json.NewDecoder(resp.Body).Decode(&respBody201)
	default:
		err = goraml.NewAPIError(resp, nil)
	}

	return respBody201, resp, err
}
