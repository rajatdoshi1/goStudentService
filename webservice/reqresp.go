package webservice

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type (
	GetUserRequest struct {
		ID   int    `json:"id"`
		Name string `json:"Name"`
	}
	GetUserResponse struct {
		ID   int    `json:"id"`
		Name string `json:"Name"`
	}
)

type GetStudentsResponses []GetUserResponse

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeNameReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req GetUserRequest
	vars := mux.Vars(r)

	req = GetUserRequest{
		Name: vars["name"],
	}
	return req, nil
}

func decodeIdReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req GetUserRequest
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	req = GetUserRequest{
		ID: id,
	}
	return req, nil
}
