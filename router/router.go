package router

import (
	controller "libr-simplified/controllers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/{ts}", controller.Get).Methods("GET")
	r.HandleFunc("/", controller.Post).Methods("POST")

	return r
}
