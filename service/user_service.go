package service

import (
	"fmt"
	"log"
	"regexp"
	"time"

	"repo"

	"github.com/bwmarrin/snowflake"
	"github.com/dgrijalva/jwt-go"
	"github.com/memcachier/bcrypt"
)

type userService struct{
	userRepo repo.AppRepository
}

type Token struct{
	jwt.StandardClaims
}

var mySigningKey []byte
var salt bcrypt.BcryptSalt = "$2a$04$pwTMbBwCyBbsuH13QnSHH."

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

func EncryptPassword(password string, salt bcrypt.BcryptSalt) (hashed string, err error) {
	hashed, err = bcrypt.Crypt(password, salt)
	if err != nil{
		fmt.Println("hash failed,",err)
	}
	return
}

func CheckPasswordHash(password, hash string) (match bool, err error){
	match, err = bcrypt.Verify(password, hash)
	return 
}

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

	userRegister.Password, err = EncryptPassword(userRegister.Password, salt)
	if err != nil {
		log.Println("Failed encrypting password", err)
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

func (s *userService) Login(email string, password string) (token string, err error) {
	mySigningKey = []byte("TheSignatureofTheBank")

	userData, err := s.userRepo.FindByEmail(email)
	if err != nil {
		fmt.Println("Error at user service, getting user data: ", err)
		return
	}

	match, err := CheckPasswordHash(password, userData.Password)
	if !match || err != nil {
		log.Println("Wrong password")
		log.Println(err)
		return
	}

	claims := Token{
		jwt.StandardClaims{
			Subject:   userData.ID,
			ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
		},
	}

	signing := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ = signing.SignedString(mySigningKey)
	if len(token) == 0 {
		log.Println("Failed to generate token")
		return
	}
	return
}

func (s *userService) ChangePassword(token string, password string, newPassword string) (success bool, err error) {
	success = false

	var id string

	at(time.Unix(0, 0), func() {
		tokenClaims, err := jwt.ParseWithClaims(token, &Token{}, func(tokenClaims *jwt.Token) (interface{}, error) {
			return []byte("IDKWhatThisIs"), nil
		})

		if claims, _ := tokenClaims.Claims.(*Token); claims.ExpiresAt > time.Now().Unix() {
			id = claims.StandardClaims.Subject
			log.Println(claims.Subject)
		} else {
			fmt.Println("token Invalid,    ", err)
		}
	})

	userData, err := s.userRepo.FindByID(id)
	if err != nil {
		fmt.Println("Error at user service: ", err)
		return
	}

	match,err := CheckPasswordHash(password, userData.Password)
	if !match || err != nil {
		log.Println("Wrong password")
		log.Println(err)
		return
	}

	hashedNewPass, err := EncryptPassword(newPassword, salt)
	if err != nil {
		log.Println("Failed encrypting password,  ", err)
		return
	}

	success, err = s.userRepo.UpdatePassword(id, hashedNewPass)
	if err != nil {
		log.Println("Error at user service, updating password: ", err)
		return
	}

	return
}

func (s *userService) GetGenres() (genres []repo.GenreList, err error){
	genres, err = s.userRepo.GetGenres()
	return
}

func (s *userService) GetInstruments() (instruments []repo.InstrumentList, err error){
	instruments, err = s.userRepo.GetInstruments()
	return
}