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
	route.HandleFunc("/getgenres", handler.GetGenresHandler).Methods("GET")
	route.HandleFunc("/getinstruments", handler.GetInstrumentsHandler).Methods("GET")
	return route
}