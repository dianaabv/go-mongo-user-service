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
	"go.mongodb.org/mongo-driver/bson/primitive" // for BSON ObjectID
	"gokit-example/account/config"
)


var RepoErr = errors.New("Unable to handle Repo Request")

const (
	database   = "buddyApp"
	collectionUsers = "sys_users"
	collectionTokens = "sys_tokens"
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

func (repo *repo) CreateUser(ctx context.Context, user User) (bool, string, User, error) {
	// for env varibales
	conf := config.New()
	collection := repo.db.Database(database).Collection(collectionUsers)
	pwd := helpers.HashAndSalt([]byte(user.Password))
	user.Password = pwd
	insertResult, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		fmt.Printf("error type: %T", err)
		// most likely an email is already registered
		// return err
		return false, "E-mail is already in use", user, nil
	}
	tk := &Token{}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, error := token.SignedString([]byte(conf.Jwtsecret.SecretKey))
	if error != nil {
		return false, "Could not create a token", user, nil
	}
	expiresAt := time.Now().Add(time.Minute * 100000).Unix()
	tk = &Token{
		// ID: user.ID,
		Email:  user.Email,
		Token: tokenString,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}	
	collectionToken := repo.db.Database(database).Collection(collectionTokens)
	insertTokenResult, err := collectionToken.InsertOne(context.TODO(), tk)
	if err != nil {
		return false, "Could not save a token", user, nil
	}
	// Vevery bad temporary solution, emailcenter requires separate microservice and queue for sending
	email := helpers.MailCenter(user.Email, user.Name, tokenString)
	fmt.Println(email, "email")
    fmt.Println("Inserted a Single Document: ", insertResult.InsertedID, insertTokenResult.InsertedID)
	return true, "User Created", user, nil
}

func (repo *repo) UpdateUser(ctx context.Context, id string, user User) (string, error) {
	if user.Email == "" || user.Password == "" {
		return "Some info is missing", nil
	}
	filter := bson.M{
		"id": id,
	}
	update := bson.M{"$set": bson.M{"email": user.Email}}
	collection := repo.db.Database(database).Collection(collectionUsers)
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
		return "User Updated", nil
	} else {
		return "User Not Found", nil
	}
	// return "", nil
}

func (repo *repo) GetUser(ctx context.Context, id string) (bool, string, User, error) {
	var user User
	fmt.Println("id", id)
	docID, _err := primitive.ObjectIDFromHex(id)
	if _err != nil {
		return false, "Wrong User id", user, nil
	}
	filter := bson.M{
		"_id": docID,
	}
	collection := repo.db.Database(database).Collection(collectionUsers)
	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		// Email not found
		// RepoErr difficult to handle
		// return "", "User not found", err 
		return false, "User not found", user, nil
	}
	return true, "User found", user, nil
}

func (repo *repo) DeleteUser(ctx context.Context, id string) (string, error) {
	docID, _err := primitive.ObjectIDFromHex(id)
	if _err != nil {
		return "Wrong User id", nil
	}
	filter := bson.M{
		"_id": docID,
	}
	collection := repo.db.Database(database).Collection(collectionUsers)
	res, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		// Email not found
		return "", RepoErr
	}
	fmt.Println("res", res)
	return "Success", nil
}

func (repo *repo) GetUserLogin(ctx context.Context, email string, password string) (string, string, User, bool, error) {
	var user User
	conf := config.New()
	// if email == "" || password == "" {
	// 	return email, "", RepoErr
	// }
	filter := bson.M{
		"email": email,
	}
	collection := repo.db.Database(database).Collection(collectionUsers)
	err := collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		// Email not found
		return email, "", user, false, nil
	}
	pwdMatch := helpers.ComparePasswords(user.Password, []byte(password))
	if pwdMatch == false {
		// Invalid login credentials. Please try again
		return email, " Invalid login credentials. Please try again", user, false, nil
	}
	// fmt.Println(user.Activated)
	if (!user.Activated) {
		return email, "Activate your account please", user, false, nil
		// fmt.Println(user.Activated, "activate your akk")
	}
	expiresAt := time.Now().Add(time.Minute * 100000).Unix()

	tk := &User{
		// ID: user.ID,
		Email:  user.Email,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, error := token.SignedString([]byte(conf.Jwtsecret.SecretKey))
	if error != nil {
		// something went wrong
		return email, "", user, false, nil
	}
	fmt.Println("tokenString", tokenString)
	return email, tokenString, user, true, nil
}
