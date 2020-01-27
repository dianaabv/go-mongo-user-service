package account

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	// "github.com/gofrs/uuid"
	"fmt"
)

type service struct {
	repostory Repository
	logger    log.Logger
}

func NewService(rep Repository, logger log.Logger) Service {
	return &service{
		repostory: rep,
		logger   : logger,
	}
}

// func (s service) CreateUser(ctx context.Context, email string, password string, name string, lastname string, phone string, dob string, country string, bio string, activated bool) (string, error) {
func (s service) CreateUser(ctx context.Context, user User) (string, error) {
	logger := log.With(s.logger, "method", "CreateUser")
	// uuid, _ := uuid.NewV4()
	// id := uuid.String()
	// user := User{
	// 	ID:       id,
	// 	Email:    email,
	// 	Password: password,
	// }
	message, err := s.repostory.CreateUser(ctx, user);
	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	// logger.Log("create user", id)

	return message, nil
}
// TODO
func (s service) UpdateUser(ctx context.Context, id string, email string, password string) (string, error) {
	logger := log.With(s.logger, "method", "UpdateUser")
	fmt.Println(id, "im hereeeeeeeeeeeee")
	user := User{
		// ID:       id,
		Email:    email,
		Password: password,
	}
	message, err := s.repostory.UpdateUser(ctx, id, user)
	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	// logger.Log("Get user", id)

	return message, nil
}

func (s service) GetUser(ctx context.Context, id string) (bool, string, User, error) {
	logger := log.With(s.logger, "method", "GetUser")

	ok, message, user, err := s.repostory.GetUser(ctx, id)
	if err != nil {
		// fmt.Println(email, message, err, "message")
		level.Error(logger).Log("err", err)
		return ok, message, user, err
		// return "", message, err
	}

	logger.Log("Get user", id)

	return ok, message, user, nil
}

func (s service) DeleteUser(ctx context.Context, id string) (string, error) {
	logger := log.With(s.logger, "method", "DeleteUser")
	res, err := s.repostory.DeleteUser(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("Delete user", id)

	return res, nil
}

func (s service) GetUserLogin(ctx context.Context, email string, password string) (string, string, User, bool, error) {
	logger := log.With(s.logger, "method", "GetUserLogin")

	email, token, user, ok, err := s.repostory.GetUserLogin(ctx, email, password)

	if err != nil {
		level.Error(logger).Log("err", err)
		return "", "", user, ok, err
	}

	return email, token, user, ok, nil
}