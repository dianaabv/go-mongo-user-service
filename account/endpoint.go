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
	VerifyUser endpoint.Endpoint
	RepeatVerifyUser endpoint.Endpoint
	ForgotPassword endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		CreateUser : makeCreateUserEndpoint(s),
		GetUser    : makeGetUserEndpoint(s),
		UserLogin  : makeGetUserLoginEndpoint(s),
		DeleteUser : makeDeleteUserEndpoint(s),
		UpdateUser : makeUpdateUserEndpoint(s),
		VerifyUser : makeVerifyUserEndpoint(s),
		RepeatVerifyUser : makeRepeatVerifyUser(s),
		ForgotPassword: makeForgotPassword(s),
	}
}

func makeCreateUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateUserRequest)
		if (req.Email == "" || req.Password == "") {
			return CreateUserResponse{Ok: false, Message: "Some data is missing"}, nil
		}
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
		ok, message, user, err := s.CreateUser(ctx, user)
		return CreateUserResponse{Ok: ok, Message: message, Respuser: user,}, err
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

func makeVerifyUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(VerifyUserRequest)
		ok, err := s.VerifyUser(ctx, req.Token, req.Email)
		fmt.Println(err)
		return VerifyUserResponse{
			Ok: ok,
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

func makeRepeatVerifyUser(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RepeatVerifyUserRequest)
		// fmt.Println(req)
		ok, message, err := s.RepeatVerifyUser(ctx, req.Email)

		return RepeatVerifyUserResponse{
			Ok		: ok,
			Message	: message,
		}, err
	}
}
func makeForgotPassword(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ForgotPasswordRequest)
		// fmt.Println(req)
		ok, message, err := s.RepeatVerifyUser(ctx, req.Email)
		return RepeatVerifyUserResponse{
			Ok		: ok,
			Message	: message,
		}, err
	}

}