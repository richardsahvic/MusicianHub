package route

import (
	"handler"

	"github.com/gorilla/mux"
)

func Routes() *mux.Router{
	route := mux.NewRouter()
	route.HandleFunc("/register", handler.RegisterHandler).Methods("POST")
	route.HandleFunc("/login", handler.LoginHandler).Methods("POST")
	route.HandleFunc("/changepassword", handler.ChangePasswordHandler).Methods("POST")
	return route
}