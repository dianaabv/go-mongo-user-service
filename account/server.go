package account

import (
	// "fmt"
	// "strings"
	"context"
	"net/http"
	// "encoding/json"
	"github.com/gorilla/mux"
	"gokit-example/account/models"
	"gokit-example/account/helpers"
	httptransport "github.com/go-kit/kit/transport/http"
)
// Exception struct
type Exception models.Exception


func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	r.Methods("POST").Path("/account/v1/user").Handler(httptransport.NewServer(
		endpoints.CreateUser,
		decodeUserReq,
		encodeResponse,
	))

	// r.Methods("GET").Path("/user/{id}").Handler(httptransport.NewServer(
	// 	endpoints.GetUser,
	// 	decodeEmailReq,
	// 	encodeResponse,
	// ))

	r.Methods("POST").Path("/account/v1/login").Handler(httptransport.NewServer(
		endpoints.UserLogin,
		decodeUserLoginReq,
		encodeResponse,
	))
	r.Methods("GET").Path("/account/v1/user/{id}").Handler(httptransport.NewServer(
		endpoints.GetUser,
		decodeEmailReq,
		encodeResponse,
	))
	r.Methods("GET").Path("/account/v1/verify/{token}/{email}").Handler(httptransport.NewServer(
		endpoints.VerifyUser,
		decodeVerifyUserReq,
		encodeResponse,
	))
	r.Methods("GET").Path("/account/v1/repeatverify/{email}").Handler(httptransport.NewServer(
		endpoints.RepeatVerifyUser,
		decodeRepeatVerifyReq,
		encodeResponse,
	))
	// Auth route
	s := r.PathPrefix("/account/v1/auth").Subrouter()
	s.Use(helpers.JwtVerify)
	
	s.Methods("GET").Path("/user/{id}").Handler(httptransport.NewServer(
		endpoints.GetUser,
		decodeEmailReq,
		encodeResponse,
	))

	s.Methods("DELETE").Path("/delete/{id}").Handler(httptransport.NewServer(
		endpoints.DeleteUser,
		decodeEmailReq,
		encodeResponse,
	))

	s.Methods("PATCH").Path("/update/{id}").Handler(httptransport.NewServer(
		endpoints.UpdateUser,
		decodeUpdateUserReq,
		encodeResponse,
	))

	return r

}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Content-Type", "*")
		next.ServeHTTP(w, r)
	})
}
