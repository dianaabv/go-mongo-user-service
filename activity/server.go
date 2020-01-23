package activity

import (
	// "fmt"
	// "strings"
	"context"
	"net/http"
	"github.com/gorilla/mux"
	httptransport "github.com/go-kit/kit/transport/http"
	// "encoding/json"
	// "gokit-example/account/models"
	"gokit-example/account/helpers"
)


func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)
	// Auth route
	s := r.PathPrefix("/activity/v1/auth").Subrouter()
	s.Use(helpers.JwtVerify)
		

	s.Methods("POST").Path("/createactivity").Handler(httptransport.NewServer(
		endpoints.CreateActivity,
		decodeUserReq,
		encodeResponse,
	))

	// r.Methods("GET").Path("/user/{id}").Handler(httptransport.NewServer(
	// 	endpoints.GetUser,
	// 	decodeEmailReq,
	// 	encodeResponse,
	// ))

	// r.Methods("POST").Path("/login").Handler(httptransport.NewServer(
	// 	endpoints.UserLogin,
	// 	decodeUserLoginReq,
	// 	encodeResponse,
	// ))
	
	// // Auth route
	// s := r.PathPrefix("/auth").Subrouter()
	// s.Use(JwtVerify)
	// s.Methods("GET").Path("/user/{id}").Handler(httptransport.NewServer(
	// 	endpoints.GetUser,
	// 	decodeEmailReq,
	// 	encodeResponse,
	// ))

	return r

}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
