package webservice

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetStudentById    endpoint.Endpoint
	GetStudentsByName endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		GetStudentById:    makeGetStudentByIdEndpoint(s),
		GetStudentsByName: makeGetStudentByNameEndpoint(s),
	}
}

func makeGetStudentByIdEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserRequest)
		student, err := s.GetStudentById(ctx, req.ID)
		return student, err
	}
}

func makeGetStudentByNameEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserRequest)
		students, err := s.GetStudentsByName(ctx, req.Name)

		return students, err
	}
}
