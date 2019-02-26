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
	route.HandleFunc("/makeprofile", handler.MakeProfileHandler).Methods("POST")
	route.HandleFunc("/updateprofile", handler.UpdateProfileHandler).Methods("POST")
	route.HandleFunc("/newpost", handler.NewPostHandler).Methods("POST")
	route.HandleFunc("/deletepost", handler.DeletePostHandler).Methods("POST")
	route.HandleFunc("/followuser", handler.FollowUserHandler).Methods("POST")
	return route
}