package webservice

import (
	"context"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	r.Methods("GET").Path("/getstudentbyid/{id}").Handler(httptransport.NewServer(
		endpoints.GetStudentById,
		decodeIdReq,
		encodeResponse,
	))

	r.Methods("GET").Path("/getstudentsbyname/{name}").Handler(httptransport.NewServer(
		endpoints.GetStudentsByName,
		decodeNameReq,
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
