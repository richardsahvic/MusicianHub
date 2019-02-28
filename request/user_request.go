package request

import (
	"repo"
)

type Response struct {
	Message string `json:"message"`
}

type RegisterRequest struct{
	ID 			string `json:"user_id"`
	Email 		string `json:"email"`
	Password 	string `json:"password"`
}

type LoginRequest struct {
	Email 		string `json:"email"`
	Password 	string `json:"password"`
}

type ChangePasswordRequest struct {
	Password    string `json:"password"`
	NewPassword string `json:"newpassword"`
}

type ProfileRequest struct {
	Name 		string 	`json:"name"`
	Gender 		string 	`json:"gender"`
	Birthdate 	string 	`json:"birthdate"`
	Bio 		string 	`json:"bio"`
	AvatarUrl	string	`json:"avatar_url"`
	Genre 		string	`json:"genre_id"`
	Instrument 	string	`json:"instrument_id"`
}

type NewPostRequest struct {
	PostId 		string `json:"post_id"`
	UserId 		string `json:"user_id"`
	PostType 	string `json:"post_type"`
	FileUrl 	string `json:"file_url"`
	Caption 	string `json:"caption"`
}

type DeletePostRequest struct {
	PostId string `json:"post_id"`
}

type FollowRequest struct {
	UserId 		string `json:"user_id"`
	FollowedId 	string `json:"followed_id"`
}

type ViewProfileResponse struct{
	Profile 	repo.UserDetail `json:"profile"`
	Following 	[]repo.UserDetail `json:"following"`
	Follower 	[]repo.UserDetail `json:"follower"`
	Posts		[]repo.UserPost `json:"posts"`
}