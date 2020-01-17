package activity

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/bson"
	"errors"
	"fmt"
	"github.com/go-kit/kit/log"
	// "github.com/dgrijalva/jwt-go"
	// "gokit-example/account/helpers"
	// "time"
)

var RepoErr = errors.New("Unable to handle Repo Request")
const (
	database   = "buddyApp"
	collection = "goActivities"
)
type repo struct {
	db     *mongo.Client
	logger log.Logger
}

func NewRepo(db *mongo.Client, logger log.Logger) Repository {
	return &repo{
		db:     db,
		logger: log.With(logger, "repo", "mongodb"),
	}
}


func (repo *repo) CreateActivity(ctx context.Context, activity Activity) error {
	// if user.Email == "" || user.Password == "" {
	// 	return RepoErr
	// }
	collection := repo.db.Database(database).Collection(collection)
	// pwd := helpers.HashAndSalt([]byte(user.Password))
	// user.Password = pwd
	fmt.Println("activity", activity)
	insertResult, err := collection.InsertOne(context.TODO(), activity)
	if err != nil {
        return err
    }
    fmt.Println("Inserted a Single Document: ", insertResult.InsertedID)
	return nil
}