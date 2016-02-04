package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

type GistsInterface interface {
}

func GistsInterfaceRoutes(r *mux.Router, i GistsInterface) {

}
