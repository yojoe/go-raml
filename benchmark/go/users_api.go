package main

import (
	"encoding/json"
	"net/http"
)

// UsersAPI is API implementation of /users root endpoint
type UsersAPI struct {
}

// Get is the handler for GET /users
// Get random user
func (api UsersAPI) Get(w http.ResponseWriter, r *http.Request) {
	respBody := User{Name: "John", Username: "Doe"}
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// Post is the handler for POST /users
// Add user
func (api UsersAPI) Post(w http.ResponseWriter, r *http.Request) {
	var reqBody User

	// decode request
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(400)
		return
	}

	// validate request
	/*if err := reqBody.Validate(); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
		return
	}*/
	respBody := reqBody.Name
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// usernameGet is the handler for GET /users/{username}
// Get information on a specific user
func (api UsersAPI) usernameGet(w http.ResponseWriter, r *http.Request) {
	var respBody User
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}
