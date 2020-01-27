package models
import (
	jwt "github.com/dgrijalva/jwt-go"
)

type Token struct {
	// ID        string `json:"id,omitempty"`
	Email     string `json:"email"`
	Token 	  string `json:"token"`
	*jwt.StandardClaims
}