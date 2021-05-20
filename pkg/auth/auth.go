package auth

import "github.com/dgrijalva/jwt-go"

type JwtCustomClaims struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	ID        string `json:"_id"`
	jwt.StandardClaims
}
