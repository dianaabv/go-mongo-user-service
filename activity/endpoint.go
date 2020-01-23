package activity

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"fmt"
)


type Endpoints struct {
	CreateActivity endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		CreateActivity: makeCreateActivityEndpoint(s),
	}
}

func makeCreateActivityEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateActivityRequest)
		fmt.Println(req)
		message, ok, err := s.CreateActivity(ctx, req.Name, req.Location)
		fmt.Println(req, ok)
		return CreateActivityResponse{Message: message, Ok: ok}, err
	}
}