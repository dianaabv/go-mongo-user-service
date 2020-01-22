package account

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"fmt"
)

type Endpoints struct {
	CreateUser endpoint.Endpoint
	GetUser    endpoint.Endpoint
	UserLogin  endpoint.Endpoint
	DeleteUser endpoint.Endpoint
	UpdateUser endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		CreateUser : makeCreateUserEndpoint(s),
		GetUser    : makeGetUserEndpoint(s),
		UserLogin  : makeGetUserLoginEndpoint(s),
		DeleteUser : makeDeleteUserEndpoint(s),
		UpdateUser : makeUpdateUserEndpoint(s),
	}
}

func makeCreateUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateUserRequest)
		ok, err := s.CreateUser(ctx, req.Email, req.Password)
		return CreateUserResponse{Ok: ok}, err
	}
}

func makeUpdateUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateUserRequest)
		ok, err := s.UpdateUser(ctx, req.Id, req.Email,req.Password)
		return UpdateUserResponse{Ok: ok}, err
	}
}

func makeGetUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserRequest)
		email, message, err := s.GetUser(ctx, req.Id)
		return GetUserResponse{
			Email: email,
			Message: message,
		}, err
	}
}

func makeDeleteUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserRequest)
		Ok, err := s.DeleteUser(ctx, req.Id)
		fmt.Println(Ok, "ok")
		return DeleteUserResponse{
			Ok: Ok,
		}, err
	}
}

func makeGetUserLoginEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateUserLoginRequest)
		// fmt.Println(req)
		email, token, err := s.GetUserLogin(ctx, req.Email, req.Password)

		return CreateUserLoginResponse{
			Email   : email,
			Token   : token,
		}, err
	}
}
