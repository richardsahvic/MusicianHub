package service

import (
	"repo"
)

type UserService interface {
	Register(userRegister repo.UserDetail) (bool, error)
	Login(username string, password string) (string, error)
	ChangePassword(token string, password string, newPassword string) (bool, error)
	AllGenre() (struct)
}

var User UserService

func NewService(userRepo repo.AppRepository) {
	User = NewUserService(userRepo)
}