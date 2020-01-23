package activity

import "context"

type Service interface {
	CreateActivity(ctx context.Context, name string, location string) (string, bool, error)
}

