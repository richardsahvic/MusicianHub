package service

import (
	"repo"
)

type UserService interface {
	Register(userRegister repo.UserDetail) (bool, error)
	Login(email string, password string) (string, error)
	ChangePassword(token string, password string, newPassword string) (bool, error)
	GetGenres() ([]repo.GenreList, error)
	GetInstruments() ([]repo.InstrumentList, error)
	MakeProfile(token string, profile repo.UserDetail, genre repo.UserGenre, instrument repo.UserInstrument) (bool, error)
	UpdateProfile(token string, profile repo.UserDetail, genre string, instrument string) (bool, error)
	UploadNewPost(token string, newPost repo.UserPost) (bool, error)
	DeletePost(postId string) (bool, error)
	FollowUser(token string, userFollow repo.UserFollow) (bool, error)
	ViewProfile(token string) (repo.UserDetail, []repo.UserDetail, []repo.UserDetail, []repo.UserPost, error)
	Timeline(token string) ([]repo.UserPost, error)
}

var User UserService

func NewService(userRepo repo.AppRepository) {
	User = NewUserService(userRepo)
}