package account

import (
	"context"
	// "gokit-example/account/models"
)

// type User models.User
type Repository interface {
	CreateUser(ctx context.Context, user User) (string, error)
	GetUser(ctx context.Context, id string) (bool, string, User, error)
	DeleteUser(cts context.Context, id string) (string, error) 
	GetUserLogin(ctx context.Context, email string, password string) (string, string, User, bool, error)
	UpdateUser(ctx context.Context, id string, user User) (string, error)
}