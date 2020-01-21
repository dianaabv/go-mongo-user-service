package activity

import "context"

type Service interface {
	CreateActivity(ctx context.Context, id string) (string, error)
}

