package account

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
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
		Email string `json:"email"`
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
	vars := mux.Vars(r)
	req = CreateUserLoginRequest{
		Email: vars["email"],
		Password: vars["password"],
	}
    return req, nil
}