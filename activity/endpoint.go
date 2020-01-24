package activity

import (
	"context"
	// "encoding/hex" // hexadecimal encoding of BSON obj
	"github.com/go-kit/kit/endpoint"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive" // for BSON ObjectID

)


type Endpoints struct {
	CreateActivity endpoint.Endpoint
	GetActivity    endpoint.Endpoint
	DeleteActivity endpoint.Endpoint
	UpdateActivity endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		CreateActivity : makeCreateActivityEndpoint(s),
		GetActivity    : makeGetActivityEndpoint(s),
		DeleteActivity : makeDeleteActivityEndpoint(s),
		UpdateActivity : makeUpdateActivityEndpoint(s),
	}
}

func makeCreateActivityEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateActivityRequest)
		fmt.Println(req)
		docID, err := primitive.ObjectIDFromHex(req.Owner)
		activity := Activity{
			Name        :    req.Name,
			Category    :    req.Category,
			Location    :    req.Location,
			Date        :    req.Date,
			Maxpeople   :    req.Maxpeople,
			// Photo 	  string `json:"photo"`
			Description :    req.Description,
			Owner       :    docID,
			Placesleft  :    req.Placesleft,
		}
		message, ok, err := s.CreateActivity(ctx, activity)
		fmt.Println(req, ok)
		return CreateActivityResponse{
			Ok: ok,
			Message: message,
		}, err
	}
}

func makeGetActivityEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetActivityRequest)
		name, message, ok, err := s.GetActivity(ctx, req.Id)
		return GetActivityResponse{
			Name: name,
			Message: message,
			Ok: ok,
		}, err
	}
}

func makeDeleteActivityEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetActivityRequest)
		Message, Ok, err := s.DeleteActivity(ctx, req.Id)
		fmt.Println(Ok, Message, "ok")
		return DeleteActivityResponse{
			Ok: Ok,
		}, err
	}
}


func makeUpdateActivityEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateActivityRequest)
		message, ok, err := s.UpdateActivity(ctx, req.Id, req.Name,req.Location)
		fmt.Println(message, "message")
		return UpdateActivityResponse{
			Ok: ok,
			Message: message,
		}, err
	}
}

