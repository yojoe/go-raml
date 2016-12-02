package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	rootURL = "http://localhost:5000"
)

type goramldir struct {
	client     http.Client
	AuthHeader string // Authorization header, will be sent on each request if not empty
}

func Newgoramldir() *goramldir {
	c := new(goramldir)
	c.client = http.Client{}
	return c
}

// Get list of all developers
func (c *goramldir) UsersGet(headers, queryParams map[string]interface{}) ([]User, *http.Response, error) {
	qsParam := buildQueryString(queryParams)
	var u []User

	// create request object
	req, err := http.NewRequest("GET", rootURL+"/users"+qsParam, nil)
	if err != nil {
		return u, nil, err
	}

	if c.AuthHeader != "" {
		req.Header.Set("Authorization", c.AuthHeader)
	}

	for k, v := range headers {
		req.Header.Set(k, fmt.Sprintf("%v", v))
	}

	//do the request
	resp, err := c.client.Do(req)
	if err != nil {
		return u, nil, err
	}
	defer resp.Body.Close()

	return u, resp, json.NewDecoder(resp.Body).Decode(&u)
}

// Add user
func (c *goramldir) UsersPost(user User, headers, queryParams map[string]interface{}) (User, *http.Response, error) {
	qsParam := buildQueryString(queryParams)
	var u User

	resp, err := c.doReqWithBody("POST", rootURL+"/users", &user, headers, qsParam)
	if err != nil {
		return u, nil, err
	}
	defer resp.Body.Close()

	return u, resp, json.NewDecoder(resp.Body).Decode(&u)
}

// Get information on a specific user
func (c *goramldir) UsersUsernameGet(username string, headers, queryParams map[string]interface{}) (User, *http.Response, error) {
	qsParam := buildQueryString(queryParams)
	var u User

	// create request object
	req, err := http.NewRequest("GET", rootURL+"/users/"+username+qsParam, nil)
	if err != nil {
		return u, nil, err
	}

	if c.AuthHeader != "" {
		req.Header.Set("Authorization", c.AuthHeader)
	}

	for k, v := range headers {
		req.Header.Set(k, fmt.Sprintf("%v", v))
	}

	//do the request
	resp, err := c.client.Do(req)
	if err != nil {
		return u, nil, err
	}
	defer resp.Body.Close()

	return u, resp, json.NewDecoder(resp.Body).Decode(&u)
}
