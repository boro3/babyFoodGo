package password

import (
	"golang.org/x/crypto/bcrypt"
)

//HashPassword generates hashed string for given string as input.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 15)
	return string(bytes), err
}

//CheckPasswordHash compares hashed and non hased strings given as input.
//Returns true if they match and false if they dont
func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
