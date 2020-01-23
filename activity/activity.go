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
	CreateActivity(ctx context.Context, activity Activity)  (string, bool, error)
	GetActivity(ctx context.Context, id string) (string, string, error)
	DeleteActivity(cts context.Context, id string) (string, error) 
	UpdateActivity(ctx context.Context, id string, activity Activity) (string, error)
}