package activity

import (
	// "fmt"
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gofrs/uuid"
)

type service struct {
	repostory Repository
	logger    log.Logger
}

func NewService(rep Repository, logger log.Logger) Service {
	return &service{
		repostory: rep,
		logger:    logger,
	}
}
func (s service) CreateActivity(ctx context.Context, name string, location string) (string, bool, error) {
	logger := log.With(s.logger, "method", "CreateActivity")
	uuid, _ := uuid.NewV4()
	id := uuid.String()
	var activity Activity
	// level.Error(logger).Log("err", err)
	activity = Activity{
		ID:       id,
		Name:     name,
		Location: location,
	}
	message, ok, err := s.repostory.CreateActivity(ctx, activity);
	if  err != nil {
		level.Error(logger).Log("err", err)
		return "", false, nil
	}

	logger.Log("create activity",  "done")

	return message, ok, nil
	// return "", nil

}