package activity

import (
	"context"
	"encoding/json"
	"net/http"

	// "github.com/gorilla/mux"
	"fmt"
)

type (
	CreateActivityRequest struct {
		ID       string `json:"id,omitempty"`
		Name     string `json:"name"`
		Location string `json:"location"`
	}
	CreateActivityResponse struct {
		// Email string `json:"email"`
		// Token string `json:"token"`
		Ok bool `json:"ok"`
	}
)


func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeUserReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req CreateActivityRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	fmt.Println(req)
	return req, nil
}

// func decodeEmailReq(ctx context.Context, r *http.Request) (interface{}, error) {
// 	var req GetUserRequest
// 	vars := mux.Vars(r)

// 	req = GetUserRequest{
// 		Id: vars["id"],
// 	}
// 	return req, nil
// }

// func decodeUserLoginReq(ctx context.Context, r *http.Request) (interface{}, error) {
// 	var req CreateUserLoginRequest
// 	// var res CreateUserLoginResponse
// 	// super important line
// 	err := json.NewDecoder(r.Body).Decode(&req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return req, nil
// }
