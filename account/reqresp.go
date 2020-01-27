package account

import (
	"context"
	"encoding/json"
	"net/http"
	"gokit-example/account/models"
	"github.com/gorilla/mux"
	// "reflect"
	// "fmt"
	// "io/ioutil"
)
type User models.User
type Token models.Token
// func (user User) IsStructureEmpty() bool {
// 	return reflect.DeepEqual(user, User{})
// }
type (
	CreateUserRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Name 	 string `json:"name"`
		Lastname string `json:"lastname"`
		Phone 	 string `json:"phone"`
		Dob 	 string `json:"dob"`
		Country  string `json:"country"`
		Bio 	 string `json:"bio"`
		// Photo 	 []byte `json:"photo"`
		Photo 	 string `json:"photo"`
		Activated bool   `json:"activated"`
	}
	CreateUserResponse struct {
		Ok bool `json:"ok"`
		Message string `json:"message"`
		Respuser User `json:"user"`
	}

	GetUserRequest struct {
		Id string `json:"id"`
	}
	GetUserResponse struct {
		Ok bool `json:"ok"`
		Message string `json:"message"`
		Respuser User `json:"user"`
	}
	DeleteUserRequest struct {
		Id string `json:"id"`
	}
	DeleteUserResponse struct {
		Ok string `json:"ok"`
	}
	UpdateUserRequest struct {
		Id string `json:"id"`
		// User User `json:"user"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	UpdateUserResponse struct {
		Ok string `json:"ok"`
	}
	CreateUserLoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	CreateUserLoginResponse struct {
		Email string `json:"email"`
		Token string `json:"token"`
		Respuser User `json:"user"`
		Ok 	  bool `json:"ok"`
	}
)


func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeUserReq(ctx context.Context, r *http.Request) (interface{}, error) {
	// TODO  refine it
	var req CreateUserRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	//photo file logic
	// file, handler, err := r.FormFile("photo")
    // if err != nil {
    //     fmt.Println("Error Retrieving the File")
    //     fmt.Println(err)
    //     //return
	// }
	// defer file.Close()
    // fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	// fmt.Printf("File Size: %+v\n", handler.Size)
	//  // byte array
	// fileBytes, err := ioutil.ReadAll(file)
	// if err != nil {
	//     fmt.Println(err)
	// }
	// req.Photo = fileBytes
	if err != nil {
		return nil, err
	}
	return req, nil

	// fmt.Printf(r.FormValue("bio"))
	// var req = r.ParseForm()
	// return req, nil
}
func decodeUpdateUserReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req UpdateUserRequest
	vars := mux.Vars(r)

	req = UpdateUserRequest{
		Id: vars["id"],
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeEmailReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req GetUserRequest
	vars := mux.Vars(r)

	req = GetUserRequest{
		Id: vars["id"],
	}
	return req, nil
}

func decodeUserLoginReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req CreateUserLoginRequest
	// var res CreateUserLoginResponse
	// super important line
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}
