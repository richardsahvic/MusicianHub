package service

import (
	"fmt"
	"log"
	"regexp"
	"time"

	"repo"

	"github.com/bwmarrin/snowflake"
	"github.com/dgrijalva/jwt-go"
	"github.com/jankuo/xxtea/xxtea"
)

type userService struct{
	userRepo repo.AppRepository
}

type Token struct{
	jwt.StandardClaims
}

var mySigningKey []byte

func at(t time.Time, f func()) {
	jwt.TimeFunc = func() time.Time {
		return t
	}
	f()
	jwt.TimeFunc = time.Now
}

func NewUserService(userRepo repo.AppRepository) UserService {
	s := userService{userRepo: userRepo}
	return &s
}

func EncryptPassword(password string) (string) {
	key := "userpass"
	encrypted := xxtea.Encrypt([]byte(password), []byte(key))
	return string(encrypted)
}

// func CheckPasswordHash(password, hash string) bool {
// 	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
// 	return err == nil
// }

func (s *userService) Register(userRegister repo.UserDetail) (success bool, err error) {
	success = false

	reEmail := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	emailValid := reEmail.MatchString(userRegister.Email)
	if !emailValid {
		log.Println("Email's format is not valid.")
		return
	}

	checkEmail, err := s.userRepo.FindByEmail(userRegister.Email)
	newEmail := checkEmail.Email
	if len(newEmail) != 0 {
		success = false
		log.Printf("Email: %v is already exist", newEmail)
		return
	}

	checkUsername, err := s.userRepo.FindByUsername(userRegister.Username)
	newUsername := checkUsername.Username
	if len(newUsername) != 0 {
		success = false
		log.Printf("Username: %v is already exist", newUsername)
		return
	}

	userRegister.Password = EncryptPassword(userRegister.Password)
	if userRegister.Password == "" {
		log.Println("Failed encrypting password")
		return
	}

	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println("Failed generating snowflake id,    ", err)
		return
	}
	id := node.Generate().String()

	userRegister.ID = id

	success, err = s.userRepo.InsertNewUser(userRegister)
	if err != nil {
		fmt.Println("Error at user_service.go, ", err)
		return
	}
	return
}