package service

import (
	"repo"
)

type UserService interface {
	Register(userRegister repo.UserDetail) (bool, error)
}

var User UserService

func NewService(userRepo repo.AppRepository) {
	User = NewUserService(userRepo)
}