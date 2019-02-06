package main

import (
	"log"
	"net/http"

	"route"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	route := route.Routes()

	http.Handle("/", route)
	log.Println("SERVER STARTED")

	http.ListenAndServe(":8080", route)
}