package account

import (
	"context"
	"gokit-example/account/models"
	// jwt "github.com/dgrijalva/jwt-go"
	// "github.com/jinzhu/gorm"
)

// type User struct {
// 	ID       string `json:"id,omitempty"`
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// 	*jwt.StandardClaims
// }
type User models.User
type Repository interface {
	CreateUser(ctx context.Context, user User) error
	GetUser(ctx context.Context, id string) (string, string, error)
	DeleteUser(cts context.Context, id string) (string, error) 
	GetUserLogin(ctx context.Context, email string, password string) (string, string, error)
	UpdateUser(ctx context.Context, id string, user User) (string, error)
}