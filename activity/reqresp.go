package activity

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"fmt"
)

type (
	CreateActivityRequest struct {
		ID       string `json:"id,omitempty"`
		Name     string `json:"name"`
		Location string `json:"location"`
	}
	CreateActivityResponse struct {
		Ok bool `json:"ok"`
		Message string `json: "message"` 
	}
	GetActivityRequest struct {
		Id string `json:"id"`
	}
	GetActivityResponse struct {
		Name     string `json:"name"`
		Message  string `json:"message"`
		Ok 		 bool `json:"ok"`

	}
	DeleteActivityRequest struct {
		Id string `json:"id"`
	}
	DeleteActivityResponse struct {
		Ok 		 bool `json:"ok"`
	}
	UpdateActivityRequest struct {
		Id string `json:"id"`
		Name    string `json:"name"`
		Location string `json:"location"`

	}
	UpdateActivityResponse struct {
		Ok 		 bool `json:"ok"`
	}
)


func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeActivityReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req CreateActivityRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	fmt.Println(req)
	return req, nil
}

func decodeUpdateActivityReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req UpdateActivityRequest
	vars := mux.Vars(r)

	req = UpdateActivityRequest{
		Id: vars["id"],
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeDelGetActivityReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req GetActivityRequest
	vars := mux.Vars(r)

	req = GetActivityRequest{
		Id: vars["id"],
	}
	return req, nil
}
