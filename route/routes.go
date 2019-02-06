package route

import (
	"handler"

	"github.com/gorilla/mux"
)

func Routes() *mux.Router{
	route := mux.NewRouter()
	route.HandleFunc("/register", handler.RegisterHandler).Methods("POST")
	return route
}