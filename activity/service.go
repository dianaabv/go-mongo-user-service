package activity

import "context"

type Service interface {
	CreateActivity(ctx context.Context, activity Activity) (string, error)
}

