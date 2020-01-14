package account

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
	"errors"
	"fmt"

	"github.com/go-kit/kit/log"
	"golang.org/x/crypto/bcrypt"
)

var RepoErr = errors.New("Unable to handle Repo Request")

const (
	database   = "buddyApp"
	collection = "goUsers"
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

func hashAndSalt(pwd []byte) string {
    
    // Use GenerateFromPassword to hash & salt pwd.
    // MinCost is just an integer constant provided by the bcrypt
    // package along with DefaultCost & MaxCost. 
    // The cost can be any value you want provided it isn't lower
    // than the MinCost (4)
    hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
    if err != nil {
		fmt.Println("Inserted a Single Document: ", err)
    }
    // GenerateFromPassword returns a byte slice so we need to
    // convert the bytes to a string and return it
    return string(hash)
}

func (repo *repo) CreateUser(ctx context.Context, user User) error {
	if user.Email == "" || user.Password == "" {
		return RepoErr
	}
	collection := repo.db.Database(database).Collection(collection)
	// user.Password = hashAndSalt([]byte(user.Password))
	pwd := hashAndSalt([]byte(user.Password))
	user.Password = pwd
	insertResult, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
        return err
    }
    fmt.Println("Inserted a Single Document: ", insertResult.InsertedID)
	return nil
}

func (repo *repo) GetUser(ctx context.Context, id string) (string, error) {
	var email string
	
	// err := repo.db.QueryRow("SELECT email FROM users WHERE id=$1", id).Scan(&email)
	// if err != nil {
	// 	return "", RepoErr
	// }

	// return email, nil
	return email, nil
}
