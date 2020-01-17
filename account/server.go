package account

import (
	"fmt"
	"strings"
	"context"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"gokit-example/account/models"
	// "gokit-example/account/helpers"
	httptransport "github.com/go-kit/kit/transport/http"
	jwt "github.com/dgrijalva/jwt-go"
)
// Exception struct
type Exception models.Exception

func JwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var header = r.Header.Get("x-access-token") //Grab the token from the header

		header = strings.TrimSpace(header)
		fmt.Println(header, "nO access")
		if header == "" {
			//Token is missing, returns with error code 403 Unauthorized
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(Exception{Message: "Missing auth token"})
			return
		}
		tk := &Token{}
		_, err := jwt.ParseWithClaims(header, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})
		// fmt.Println(err, "err")
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(Exception{Message: err.Error()})
			return
		}

		ctx := context.WithValue(r.Context(), "user", tk)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	r.Methods("POST").Path("/user").Handler(httptransport.NewServer(
		endpoints.CreateUser,
		decodeUserReq,
		encodeResponse,
	))

	// r.Methods("GET").Path("/user/{id}").Handler(httptransport.NewServer(
	// 	endpoints.GetUser,
	// 	decodeEmailReq,
	// 	encodeResponse,
	// ))

	r.Methods("POST").Path("/login").Handler(httptransport.NewServer(
		endpoints.UserLogin,
		decodeUserLoginReq,
		encodeResponse,
	))
	
	// Auth route
	s := r.PathPrefix("/auth").Subrouter()
	s.Use(JwtVerify)
	s.Methods("GET").Path("/user/{id}").Handler(httptransport.NewServer(
		endpoints.GetUser,
		decodeEmailReq,
		encodeResponse,
	))

	return r

}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
