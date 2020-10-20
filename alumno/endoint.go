package alumno

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type getStudentByIDRequest struct {
	StudentID int
}

func makeGetStudentByIdEndPoint(s Service) endpoint.Endpoint {
	//Permitir la concurrencia con Context
	getStudentByIdEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getStudentByIDRequest)
		product, err := s.GetStudentById(&req)
		if err != nil {
			panic(err)
		}
		return product, nil
	}
	return getStudentByIdEndPoint
}
