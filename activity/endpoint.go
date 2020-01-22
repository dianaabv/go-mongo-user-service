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
		ok, err := s.CreateActivity(ctx, "zsd")
		fmt.Println(req, ok)
		return CreateActivityResponse{Ok: true}, err
	}
}