package usuario

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type registerUserRequest struct {
	Usuario_id       string
	Usuario_nombre   string
	Usuario_email    string
	Usuario_password string
	Usuario_avatar   string
}

type loginUserRequest struct {
	Usuario_email    string
	Usuario_password string
}

func registerUserEndPoint(s Service) endpoint.Endpoint {
	//Permitir la concurrencia con Context
	registerUserEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(registerUserRequest)
		id_user, err := s.registrarUsuario(&req)
		return id_user, err
	}
	return registerUserEndPoint
}

func loginUserEndPoint(s Service) endpoint.Endpoint {
	//Permitir la concurrencia con Context
	loginUserEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(loginUserRequest)
		user_login, err := s.LoginUsuario(&req)
		return user_login, err
	}
	return loginUserEndPoint
}

func pruebaUserEndPoint(s Service) endpoint.Endpoint {
	loginUserEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		return map[string]string{"mensaje": "tarea creada"}, nil
	}
	return loginUserEndPoint
}
