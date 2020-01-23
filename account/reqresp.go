package account

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	// "fmt"
)

type (
	CreateUserRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	CreateUserResponse struct {
		Ok string `json:"ok"`
	}

	GetUserRequest struct {
		Id string `json:"id"`
	}
	GetUserResponse struct {
		Message string `json:"message"`
		Email string `json:"email"`
	}
	DeleteUserRequest struct {
		Id string `json:"id"`
	}
	DeleteUserResponse struct {
		Ok string `json:"ok"`
	}
	UpdateUserRequest struct {
		Id string `json:"id"`
		// User User `json:"user"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	UpdateUserResponse struct {
		Ok string `json:"ok"`
	}
	CreateUserLoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	CreateUserLoginResponse struct {
		Email string `json:"email"`
		Token string `json:"token"`
		Ok bool `json:"ok"`
	}
)


func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeUserReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func decodeUpdateUserReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req UpdateUserRequest
	vars := mux.Vars(r)

	req = UpdateUserRequest{
		Id: vars["id"],
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeEmailReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req GetUserRequest
	vars := mux.Vars(r)

	req = GetUserRequest{
		Id: vars["id"],
	}
	return req, nil
}

func decodeUserLoginReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req CreateUserLoginRequest
	// var res CreateUserLoginResponse
	// super important line
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}
