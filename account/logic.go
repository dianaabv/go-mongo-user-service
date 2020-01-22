package account

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gofrs/uuid"
	"fmt"
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

func (s service) CreateUser(ctx context.Context, email string, password string) (string, error) {
	logger := log.With(s.logger, "method", "CreateUser")
	
	uuid, _ := uuid.NewV4()
	id := uuid.String()
	user := User{
		ID:       id,
		Email:    email,
		Password: password,
	}

	if err := s.repostory.CreateUser(ctx, user); err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("create user", id)

	return "Success", nil
}

func (s service) GetUser(ctx context.Context, id string) (string, string, error) {
	logger := log.With(s.logger, "method", "GetUser")

	email, message, err := s.repostory.GetUser(ctx, id)
	fmt.Println(email, message, err, "message")
	if err != nil {
		fmt.Println("im heeeere")
		// level.Error(logger).Log("err", err)
		return "", message, err
	}

	logger.Log("Get user", id)

	return email, message, nil
}

func (s service) DeleteUser(ctx context.Context, id string) (string, error) {
	logger := log.With(s.logger, "method", "DeleteUser")

	email, err := s.repostory.DeleteUser(ctx, id)

	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("Delete user", id)

	return email, nil
}

func (s service) GetUserLogin(ctx context.Context, email string, password string) (string, string, error) {
	logger := log.With(s.logger, "method", "GetUserLogin")

	email, token, err := s.repostory.GetUserLogin(ctx, email, password)

	if err != nil {
		level.Error(logger).Log("err", err)
		return "", "", err
	}

	return email, token, nil
}