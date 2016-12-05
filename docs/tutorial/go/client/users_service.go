package main

import (
	"encoding/json"
	"net/http"
)

type UsersService service

// Get list of all developers
func (s *UsersService) UsersGet(headers, queryParams map[string]interface{}) ([]User, *http.Response, error) {
	var u []User

	resp, err := s.client.doReqNoBody("GET", s.client.BaseURI+"/users", headers, queryParams)
	if err != nil {
		return u, nil, err
	}
	defer resp.Body.Close()

	return u, resp, json.NewDecoder(resp.Body).Decode(&u)
}

// Add user
func (s *UsersService) UsersPost(user User, headers, queryParams map[string]interface{}) (User, *http.Response, error) {
	var u User

	resp, err := s.client.doReqWithBody("POST", s.client.BaseURI+"/users", &user, headers, queryParams)
	if err != nil {
		return u, nil, err
	}
	defer resp.Body.Close()

	return u, resp, json.NewDecoder(resp.Body).Decode(&u)
}

// Get information on a specific user
func (s *UsersService) UsersUsernameGet(username string, headers, queryParams map[string]interface{}) (User, *http.Response, error) {
	var u User

	resp, err := s.client.doReqNoBody("GET", s.client.BaseURI+"/users/"+username, headers, queryParams)
	if err != nil {
		return u, nil, err
	}
	defer resp.Body.Close()

	return u, resp, json.NewDecoder(resp.Body).Decode(&u)
}
