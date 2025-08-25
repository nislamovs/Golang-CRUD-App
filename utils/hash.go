package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPasswd(passwd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hashedPassword string) bool {
	//return hashedPassword == password
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
