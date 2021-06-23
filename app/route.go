package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

func route(router *mux.Router, ch *UserHandlers) {
	router.HandleFunc("/users", ch.getAll).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", ch.getBuyId).Methods(http.MethodGet)
	router.HandleFunc("/users", ch.Insert).Methods(http.MethodPost)
	router.HandleFunc("/users/{id}", ch.Update).Methods(http.MethodPut)
	router.HandleFunc("/users/{id}", ch.Delete).Methods(http.MethodDelete)
}
