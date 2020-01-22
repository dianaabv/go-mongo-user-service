package models
import (
	jwt "github.com/dgrijalva/jwt-go"
)

type User struct {
	ID       string `json:"id,omitempty"`
	Email    string `json:"email"`
	Password string `json:"password"`
	*jwt.StandardClaims
}