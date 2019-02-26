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
		Password:	regRequest.Password,
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

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, _ := ioutil.ReadAll(io.LimitReader(r.Body, 5000))

	var loginReq request.LoginRequest
	json.Unmarshal(body, &loginReq)

	loginToken, err := userService.Login(loginReq.Email, loginReq.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var loginResp request.Response

	if len(loginToken) == 0 {
		loginResp.Message = "Login failed"
		w.WriteHeader(http.StatusUnauthorized)
	} else {
		loginResp.Message = "Logged in"
		w.Header().Set("token", loginToken)
	}
	json.NewEncoder(w).Encode(loginResp)
}

func ChangePasswordHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	token := r.Header.Get("token")

	body, _ := ioutil.ReadAll(io.LimitReader(r.Body, 5000))

	var changePasswordReq request.ChangePasswordRequest
	json.Unmarshal(body, &changePasswordReq)

	success, err := userService.ChangePassword(token, changePasswordReq.Password, changePasswordReq.NewPassword)
	if err != nil {
		log.Println("Failed to register: ", err)
	}

	var changePwResp request.Response

	if !success {
		changePwResp.Message = "Failed to change password"
	} else {
		changePwResp.Message = "Password changed"
	}
	json.NewEncoder(w).Encode(changePwResp)
}

func GetGenresHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	genres, _ := userService.GetGenres()
	json.NewEncoder(w).Encode(genres)
}

func GetInstrumentsHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	instruments, _ := userService.GetInstruments()
	json.NewEncoder(w).Encode(instruments)
}

func  MakeProfileHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	token := r.Header.Get("token")

	body, _ := ioutil.ReadAll(io.LimitReader(r.Body, 5000))

	var profileReq request.ProfileRequest
	json.Unmarshal(body, &profileReq)

	userProfile := repo.UserDetail{
		Name:	profileReq.Name,
		Gender: profileReq.Gender,
		Birthdate: profileReq.Birthdate,
		Bio: profileReq.Bio,
		AvatarUrl: profileReq.AvatarUrl,
	}

	userGenre := repo.UserGenre{
		GenreId: profileReq.Genre,
	}

	userInstrument := repo.UserInstrument{
		InstrumentId: profileReq.Instrument,
	}

	success, err := userService.MakeProfile(token, userProfile, userGenre, userInstrument)
	if err != nil {
		log.Println("Failed to make profile: ", err)
	}

	var makeProfileResp request.Response

	if !success {
		makeProfileResp.Message = "Failed to make profile"
	} else {
		makeProfileResp.Message = "Profile made"
	}

	json.NewEncoder(w).Encode(makeProfileResp)
}

func  UpdateProfileHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	token := r.Header.Get("token")

	body, _ := ioutil.ReadAll(io.LimitReader(r.Body, 5000))

	var profileReq request.ProfileRequest
	json.Unmarshal(body, &profileReq)

	userProfile := repo.UserDetail{
		Name:	profileReq.Name,
		Gender: profileReq.Gender,
		Birthdate: profileReq.Birthdate,
		Bio: profileReq.Bio,
		AvatarUrl: profileReq.AvatarUrl,
	}

	userGenre := profileReq.Genre

	userInstrument := profileReq.Instrument

	success, err := userService.UpdateProfile(token, userProfile, userGenre, userInstrument)
	if err != nil {
		log.Println("Failed to update profile: ", err)
	}	

	var updateProfileResp request.Response

	if !success {
		updateProfileResp.Message = "Failed to update profile"
	} else {
		updateProfileResp.Message = "Profile updated"
	}

	json.NewEncoder(w).Encode(updateProfileResp)
}

func NewPostHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	token := r.Header.Get("token")

	body, _ := ioutil.ReadAll(io.LimitReader(r.Body, 5000))

	var userPostReq request.NewPostRequest
	json.Unmarshal(body, &userPostReq)

	newPost := repo.UserPost{
		PostType: userPostReq.PostType,
		FileUrl:	userPostReq.FileUrl,
		Caption:	userPostReq.Caption,
	}

	success, err := userService.UploadNewPost(token, newPost)
	if err != nil {
		log.Println("Failed to insert new post:", err)
	}

	var newPostResp request.Response

	if !success {
		newPostResp.Message = "Failed to post"
	} else {
		newPostResp.Message = "Posted"
	}

	json.NewEncoder(w).Encode(newPostResp)
}

func DeletePostHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	body, _ := ioutil.ReadAll(io.LimitReader(r.Body, 5000))

	var deletePostReq request.DeletePostRequest
	json.Unmarshal(body, &deletePostReq)

	success, _ := userService.DeletePost(deletePostReq.PostId)

	var deletePostResp request.Response

	if !success {
		deletePostResp.Message = "Failed to delete psot"
	} else {
		deletePostResp.Message = "Post deleted"
	}

	json.NewEncoder(w).Encode(deletePostResp)
}

func FollowUserHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	token := r.Header.Get("token")

	body, _ := ioutil.ReadAll(io.LimitReader(r.Body, 5000))

	var followReq request.FollowRequest
	json.Unmarshal(body, &followReq)

	userFollow := repo.UserFollow{
		UserId:	followReq.UserId,
		FollowedId: followReq.FollowedId,
	}

	success, _ := userService.FollowUser(token, userFollow)

	var followResp request.Response

	if !success {
		followResp.Message = "Failed to follow"
	} else {
		followResp.Message = "Followed"
	}

	json.NewEncoder(w).Encode(followResp)
}