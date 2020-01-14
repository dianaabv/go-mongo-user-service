package account

import "context"

type Service interface {
	CreateUser(ctx context.Context, email string, password string) (string, error)
	GetUser(ctx context.Context, id string) (string, error)
	GetUserLogin(ctx context.Context, email string, token string) (string, string, error)
}

