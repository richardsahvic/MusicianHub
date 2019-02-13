package main

import (
	"fmt"

	"github.com/memcachier/bcrypt"
)

var contoh = "bla8888@ijo"
var utksalt bcrypt.BcryptSalt = "kunciutkappricat"

func main(){
	hashed, err := bcrypt.Crypt(contoh, utksalt)
	if err != nil{
		fmt.Println("hash failed,",err)
		return
	}
	fmt.Println(hashed)
}
