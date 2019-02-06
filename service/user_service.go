package service

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"time"

	"repo"

	"github.com/bwmarrin/snowflake"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
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

func NewUserService(userRepo repo.BankRepository) UserService {
	s := userService{userRepo: userRepo}
	return &s
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (s *userService) Register(userRepo repo.UserDetail) (success bool, err error) {
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

	userRegister.Password, err = HashPassword(userRegister.Password)
	if err != nil {
		log.Println("Failed encrypting password,  ", err)
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