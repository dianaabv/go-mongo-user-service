package activity

import "context"

type Service interface {
	CreateActivity(ctx context.Context, name string, location string) (string, bool, error)
	GetActivity(ctx context.Context, id string) (string, string, bool, error)
	DeleteActivity(cts context.Context, id string) (string, bool, error) 
	// GetUserLogin(ctx context.Context, email string, password string) (string, string, error)
	// UpdateUser(ctx context.Context, id string, user User) (string, error)
	UpdateActivity(ctx context.Context, id string, name string, location string) (string, bool, error)
}

