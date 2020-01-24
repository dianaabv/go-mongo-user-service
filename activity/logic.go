package activity

import (
	// "fmt"
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	// "github.com/gofrs/uuid"
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
func (s service) CreateActivity(ctx context.Context, activity Activity) (string, bool, error) {
	logger := log.With(s.logger, "method", "CreateActivity")
	// uuid, _ := uuid.NewV4()
	// id := uuid.String()
	// var activity Activity
	// // level.Error(logger).Log("err", err)
	// activity = Activity{
	// 	// ID:       id,
	// 	Name:     name,
	// 	Location: location,
	// }
	message, ok, err := s.repostory.CreateActivity(ctx, activity);
	if  err != nil {
		level.Error(logger).Log("err", err)
		return "", false, nil
	}

	logger.Log("create activity",  "done")

	return message, ok, nil
	// return "", nil

}
func (s service) UpdateActivity(ctx context.Context, id string, name string, location string) (string, bool, error) {
	logger := log.With(s.logger, "method", "UpdateActivity")
	// fmt.Println(id, "im hereeeeeeeeeeeee")
	activity := Activity{
		// ID:       id,
		Name:    name,
		Location: location,
	}
	message, err := s.repostory.UpdateActivity(ctx, id, activity)
	if err != nil {
		level.Error(logger).Log("err", err)
		return "", false, err
	}

	// logger.Log("Get user", id)

	return message, true, nil
}

func (s service) GetActivity(ctx context.Context, id string) (string, string, bool, error) {
	logger := log.With(s.logger, "method", "GetActivity")

	email, message, err := s.repostory.GetActivity(ctx, id)
	if err != nil {
		// fmt.Println(email, message, err, "message")
		level.Error(logger).Log("err", err)
		return "", message, false, err
	}

	logger.Log("Get user", id)

	return email, message, true, nil
}

func (s service) DeleteActivity(ctx context.Context, id string) (string, bool, error) {
	logger := log.With(s.logger, "method", "DeleteActivity")
	res, err := s.repostory.DeleteActivity(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
		return "", false, err
	}

	logger.Log("Delete activity", id)

	return res, true, nil
}
