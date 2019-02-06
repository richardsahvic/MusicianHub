package handler

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"service"

	"datasource"
	"repo"
	"request"
)

var db = datasource.InitConnection()
var userService = service.NewUserService(repo.NewRepository(db))

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	body, _ := ioutil.ReadAll(io.LimitReader(r.Body, 5000))

	var regRequest request.RegisterRequest
	json.Unmarshal(body, &regRequest)

	userRegister := repo.UserDetail{
		Email:		regRequest.Email,
		Username:	regRequest.Username,
		Password:	regRequest.Password,
		Name:		regRequest.Name,
		Gender:		regRequest.Gender,
		Birthdate:	regRequest.Birthdate,
		Bio:		regRequest.Bio,
		Role:		regRequest.Role,
	}
	
	registerResult, err := userService.Register(userRegister)
	if err != nil {
		log.Println("failed to register,    ", err)
	}

	var regResponse request.Response

	if !registerResult {
		regResponse.Message = "Register failed"
	} else {
		regResponse.Message = "Register success"
	}
	json.NewEncoder(w).Encode(regResponse)
}