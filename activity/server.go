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
		decodeActivityReq,
		encodeResponse,
	))

	s.Methods("GET").Path("/activity/{id}").Handler(httptransport.NewServer(
		endpoints.GetActivity,
		decodeDelGetActivityReq,
		encodeResponse,
	))

	s.Methods("DELETE").Path("/delete/{id}").Handler(httptransport.NewServer(
		endpoints.DeleteActivity,
		decodeDelGetActivityReq,
		encodeResponse,
	))

	s.Methods("PATCH").Path("/update/{id}").Handler(httptransport.NewServer(
		endpoints.UpdateActivity,
		decodeUpdateActivityReq,
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
