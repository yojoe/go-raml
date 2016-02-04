package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

type DeliveriesInterface interface {
	Get(http.ResponseWriter, *http.Request)

	Post(http.ResponseWriter, *http.Request)

	deliveryIdGet(http.ResponseWriter, *http.Request)

	deliveryIdPatch(http.ResponseWriter, *http.Request)

	deliveryIdDelete(http.ResponseWriter, *http.Request)
}

func DeliveriesInterfaceRoutes(r *mux.Router, i DeliveriesInterface) {

	r.HandleFunc("/deliveries", i.Get).Methods("GET")

	r.HandleFunc("/deliveries", i.Post).Methods("POST")

	r.HandleFunc("/deliveries/{deliveryId}", i.deliveryIdGet).Methods("GET")

	r.HandleFunc("/deliveries/{deliveryId}", i.deliveryIdPatch).Methods("PATCH")

	r.HandleFunc("/deliveries/{deliveryId}", i.deliveryIdDelete).Methods("DELETE")

}
