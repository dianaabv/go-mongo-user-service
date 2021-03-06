package account

import (
	"context"
	// "gokit-example/account/models"
)

// type User models.User
type Repository interface {
	CreateUser(ctx context.Context, user User) (bool, string, User, error)
	GetUser(ctx context.Context, id string) (bool, string, User, error)
	DeleteUser(cts context.Context, id string) (string, error) 
	GetUserLogin(ctx context.Context, email string, password string) (string, string, User, bool, error)
	UpdateUser(ctx context.Context, id string, user User) (string, error)
	VerifyUser(ctx context.Context, token string, email string) (bool, error)
	RepeatVerifyUser(ctx context.Context, email string) (bool, string, error)
	ForgotPassword(ctx context.Context, email string) (bool, string, error)
}