package account

import (
	"context"
	jwt "github.com/dgrijalva/jwt-go"
	// "github.com/jinzhu/gorm"
)

type User struct {
	ID       string `json:"id,omitempty"`
	Email    string `json:"email"`
	Password string `json:"password"`
	*jwt.StandardClaims
}

//Token struct declaration
// type Token struct {
// 	ID string
// 	Email  string
// 	*jwt.StandardClaims
// }
type Repository interface {
	CreateUser(ctx context.Context, user User) error
	GetUser(ctx context.Context, id string) (string, error)
	GetUserLogin(ctx context.Context, email string, password string) (string, string, error)
}