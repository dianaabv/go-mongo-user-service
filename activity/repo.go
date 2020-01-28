package activity

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"errors"
	"fmt"
	"github.com/go-kit/kit/log"
)

var RepoErr = errors.New("Unable to handle Repo Request")
const (
	database   = "buddyApp"
	collection = "sys_activities"
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


func (repo *repo) CreateActivity(ctx context.Context, activity Activity) (string, bool, error) {
	// if activity.Name == "" || activity.Location == "" {
	// 	return "Some data is missing", false, RepoErr
	// }
	
	// fmt.Println(hexByte, "hexByte")
	// blah := string(hexByte)

	// activity.Owner = hexByte
	collection := repo.db.Database(database).Collection(collection)
	// pwd := helpers.HashAndSalt([]byte(user.Password))
	// user.Password = pwd
	fmt.Println("activity", activity)
	insertResult, err := collection.InsertOne(context.TODO(), activity)
	if err != nil {
        return "Something went wrong", false, err
    }
    fmt.Println("Inserted a Single Document: ", insertResult.InsertedID)
	return "Activity created", true, nil
}

func (repo *repo) UpdateActivity(ctx context.Context, id string, activity Activity) (string, error) {
	if activity.Name == "" || activity.Location == "" {
		return "Some info is missing", nil
	}
	filter := bson.M{
		"_id": id,
	}
	update := bson.M{"$set": bson.M{"name": activity.Name}}
	collection := repo.db.Database(database).Collection(collection)
	result, err := collection.UpdateOne(
        ctx,
        filter,
        update,
	)
	if err != nil {
		// Email not found
		// RepoErr difficult to handle
		return "Error", nil
	} 
	if result.MatchedCount == 1 {
		return "Activity Updated", nil
	} else {
		return "Activity Not Found", nil
	}
	// return "", nil
}

func (repo *repo) GetActivity(ctx context.Context, id string) (string, string, error) {
	var activity Activity
	fmt.Println("id", id)
	filter := bson.M{
		"id": id,
	}
	collection := repo.db.Database(database).Collection(collection)
	err := collection.FindOne(ctx, filter).Decode(&activity)
	if err != nil {
		// Email not found
		// RepoErr difficult to handle
		// return "", "User not found", err 
		return "", "Activity not found", nil
	}
	return activity.Name, "Activity found", nil
}

func (repo *repo) DeleteActivity(ctx context.Context, id string) (string, error) {
	filter := bson.M{
		"id": id,
	}
	collection := repo.db.Database(database).Collection(collection)
	res, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		// Email not found
		return "", RepoErr
	}
	fmt.Println("res", res)
	return "Success", nil
}
