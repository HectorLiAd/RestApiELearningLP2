package usuario

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type registerUserRequest struct {
	UsuarioID       string
	UsuarioNombre   string
	UsuarioEmail    string
	UsuarioPassword string
	UsuarioAvatar   string
}

type loginUserRequest struct {
	UsuarioEmail    string
	UsuarioPassword string
}

func registerUserEndPoint(s Service) endpoint.Endpoint {
	//Permitir la concurrencia con Context
	registerUserEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(registerUserRequest)
		idUser, err := s.registrarUsuario(&req)
		return idUser, err
	}
	return registerUserEndPoint
}

func loginUserEndPoint(s Service) endpoint.Endpoint {
	//Permitir la concurrencia con Context
	loginUserEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(loginUserRequest)
		userLogin, err := s.LoginUsuario(&req)
		return userLogin, err
	}
	return loginUserEndPoint
}

func pruebaUserEndPoint(s Service) endpoint.Endpoint {
	loginUserEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		return map[string]string{"mensaje": "tarea creada"}, nil
	}
	return loginUserEndPoint
}
