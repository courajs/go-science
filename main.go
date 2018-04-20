package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	pw, ok := os.LookupEnv("PW")
	if !ok {
		panic("Pass PW")
	}
	hash, ok := os.LookupEnv("PW_HASH")
	if !ok {
		panic("Pass PW_HASH")
	}

	nah := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))

	fmt.Println(nah, pw, hash)

}
