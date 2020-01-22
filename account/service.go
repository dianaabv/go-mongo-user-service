package account

import "context"

type Service interface {
	CreateUser(ctx context.Context, email string, password string) (string, error)
	GetUser(ctx context.Context, id string) (string, string, error)
	DeleteUser(cts context.Context, id string) (string, error) 
	GetUserLogin(ctx context.Context, email string, password string) (string, string, error)
	// UpdateUser(ctx context.Context, id string, user User) (string, error)
	UpdateUser(ctx context.Context, id string, email string, password string) (string, error)
}

