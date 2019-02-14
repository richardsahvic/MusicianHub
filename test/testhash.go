package main

import (
	"fmt"

	"github.com/memcachier/bcrypt"
)

var contoh = "bla8888@ijo"
var salt bcrypt.BcryptSalt = "$2a$04$pwTMbBwCyBbsuH13QnSHH."

func main(){
	hashed, err := bcrypt.Crypt(contoh, utksalt)
	if err != nil{
		fmt.Println("hash failed,",err)
		return
	}
	fmt.Println(hashed)
}
