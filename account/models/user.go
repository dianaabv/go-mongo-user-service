package models
import (
	jwt "github.com/dgrijalva/jwt-go"
)

type User struct {
	// ID        string `json:"id,omitempty"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Name 	  string `json:"name"`
	Lastname  string `json:"lastname"`
	Phone 	  string `json:"phone"`
	Dob 	  string `json:"dob"`
	Country   string `json:"country"`
	Bio 	  string `json:"bio"`
	Photo 	  string `json:"photo"`
	// Photo 	  []byte `json:"photo"`
	Activated bool   `json:"activated"`
	*jwt.StandardClaims
}