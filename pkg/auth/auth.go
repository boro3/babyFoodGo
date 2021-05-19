package auth

import "github.com/dgrijalva/jwt-go"

type JwtCustomClaims struct {
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Email      string `json:"email"`
	ID         string `json:"_id"`
	jwt.StandardClaims
}
