package main

import (
	"fmt"
	"net/http"
)

// ResourcesAPI is API implementation of /resources root endpoint
type ResourcesAPI struct {
}

// Get is the handler for GET /resources
// Get a resource
func (api ResourcesAPI) Get(w http.ResponseWriter, r *http.Request) {
	// uncomment below line to add header
	// w.Header().Set("key","value")
	fmt.Fprintf(w, "Actual implementation should return a resource")
}

// Post is the handler for POST /resources
// Post a resource
func (api ResourcesAPI) Post(w http.ResponseWriter, r *http.Request) {
	// uncomment below line to add header
	// w.Header().Set("key","value")
	fmt.Fprintf(w, "Actual implementation should post a resource")
}
