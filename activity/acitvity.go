package activity

import (
	"context"
	// jwt "github.com/dgrijalva/jwt-go"
	// "github.com/jinzhu/gorm"
)

type Activity struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name"`
	Location string `json:"location"`
	// *jwt.StandardClaims
}

type Repository interface {
	CreateActivity(ctx context.Context, activity Activity) error
	// GetUser(ctx context.Context, id string) (string, error)
	// GetUserLogin(ctx context.Context, email string, password string) (string, string, error)
}