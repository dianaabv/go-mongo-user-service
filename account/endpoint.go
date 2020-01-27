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
		fmt.Println(req)
		// TODO 
		user := User{
			// ID:       id,
			Email:    req.Email,
			Password: req.Password,
			Name: req.Name,
			Lastname: req.Lastname,
			Phone: req.Phone,
			Dob: req.Dob,
			Country: req.Country,
			Bio: req.Bio,
			Activated: req.Activated,
			Photo: req.Photo,
		}
		// ok, err := s.CreateUser(ctx, req.Email, req.Password, req.Name, req.Lastname, req.Phone, req.Dob, req.Country, req.Bio, req.Activated)
		ok, err := s.CreateUser(ctx, user)
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
		ok, message, user, err := s.GetUser(ctx, req.Id)
		fmt.Println(user, message)
		return GetUserResponse{
			Ok: ok,
			Message: message,
			Respuser: user,
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
		email, token, user, ok, err := s.GetUserLogin(ctx, req.Email, req.Password)

		return CreateUserLoginResponse{
			Email   : email,
			Token   : token,
			Respuser: user,
			Ok		: ok,
		}, err
	}
}
