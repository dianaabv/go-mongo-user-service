package account

import "context"

type Service interface {
	// CreateUser(ctx context.Context, email string, password string, name string, lastname string, phone string, dob string, country string, bio string, activated bool) (string, error)
	CreateUser(ctx context.Context, user User) (bool, string, User, error)
	GetUser(ctx context.Context, id string) (bool, string, User, error)
	DeleteUser(cts context.Context, id string) (string, error) 
	GetUserLogin(ctx context.Context, email string, password string) (string, string, User, bool, error)
	// UpdateUser(ctx context.Context, id string, user User) (string, error)
	UpdateUser(ctx context.Context, id string, email string, password string) (string, error)
	VerifyUser(ctx context.Context, token string, email string) (bool, error)
	RepeatVerifyUser(ctx context.Context, email string) (bool, string, error)
}

