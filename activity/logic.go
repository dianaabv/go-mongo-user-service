package activity

import (
	// "fmt"
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
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
func (s service) CreateActivity(ctx context.Context, id string) (string, error) {
	logger := log.With(s.logger, "method", "CreateActivity")
	var activity Activity
	// level.Error(logger).Log("err", err)
	// activity = Activity{
	// 	ID:       "id",
	// 	Name:     "name",
	// 	Location: "location",
	// }
	if err := s.repostory.CreateActivity(ctx, activity); err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("create activity",  "asds")

	return "Success", nil
	// return "", nil

}