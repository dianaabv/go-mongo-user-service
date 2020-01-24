package activity

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive" // for BSON ObjectID

	// jwt "github.com/dgrijalva/jwt-go"
	// "github.com/jinzhu/gorm"
)

type Activity struct {
	// ID       string `json:"id,omitempty"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Location string `json:"location"`
	Date 	 string `json:"date"`
	Maxpeople int `json:"maxpeople"`
	// Photo 	  string `json:"photo"`
	Description string `json:"description"`
	Owner primitive.ObjectID `json:"owner"`
	Placesleft int `json:"placesleft"`
	// *jwt.StandardClaims
}

type Repository interface {
	CreateActivity(ctx context.Context, activity Activity)  (string, bool, error)
	GetActivity(ctx context.Context, id string) (string, string, error)
	DeleteActivity(cts context.Context, id string) (string, error) 
	UpdateActivity(ctx context.Context, id string, activity Activity) (string, error)
}