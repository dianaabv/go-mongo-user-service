package account

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"errors"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/dgrijalva/jwt-go"
	"gokit-example/account/helpers"
	"time"
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


func (repo *repo) CreateUser(ctx context.Context, user User) error {
	if user.Email == "" || user.Password == "" {
		return RepoErr
	}
	collection := repo.db.Database(database).Collection(collection)
	pwd := helpers.HashAndSalt([]byte(user.Password))
	user.Password = pwd
	fmt.Println("user", user)
	insertResult, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
        return err
    }
    fmt.Println("Inserted a Single Document: ", insertResult.InsertedID)
	return nil
}

func (repo *repo) GetUser(ctx context.Context, id string) (string, error) {
	var user User
	fmt.Println("id", id)
	filter := bson.M{
		"id": id,
	}
	collection := repo.db.Database(database).Collection(collection)
	err := collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		// Email not found
		return "", RepoErr
	}
	return user.Email, nil
}


func (repo *repo) GetUserLogin(ctx context.Context, email string, password string) (string, string, error) {
	var user User
	// var token string
	if email == "" || password == "" {
		return email, "", RepoErr
	}
	filter := bson.M{
		"email": email,
	}
	collection := repo.db.Database(database).Collection(collection)
	err := collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		// Email not found
		return email, "", nil
	}
	pwdMatch := helpers.ComparePasswords(user.Password, []byte(password))
	if pwdMatch == false {
		// Invalid login credentials. Please try again
		return email, "", nil
	}
	expiresAt := time.Now().Add(time.Minute * 100000).Unix()

	tk := &Token{
		ID: user.ID,
		Email:  user.Email,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	// TODO env variable
	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		// something went wrong
		return email, "", nil
	}

	fmt.Println("tokenString", tokenString)
    // succesful login
	return email, tokenString, nil
}
