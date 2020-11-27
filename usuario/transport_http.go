package usuario

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
)

/*MakeHTTPSHandler para poder usar los diferentes metodos de inciicio secion*/
func MakeHTTPSHandler(s Service) http.Handler {
	//Creacion de las rutas
	r := chi.NewRouter() //Creando instancia para iniciar el ruteo

	createUseHandler := kithttp.NewServer(
		registerUserEndPoint(s),
		registerUserRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodPost, "/register", createUseHandler)

	loginUseHandler := kithttp.NewServer(
		loginUserEndPoint(s),
		loginUserRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodPost, "/login", loginUseHandler)
	/*Pruebasssss*/
	pruebaUseHandler := kithttp.NewServer(
		pruebaUserEndPoint(s),
		pruebaUserRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodGet, "/verPerfil", pruebaUseHandler)

	return r
}

func registerUserRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := registerUserRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	return request, err
}

func loginUserRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := loginUserRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	return request, err
}

/*----------------------------------------*/
func pruebaUserRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := loginUserRequest{}
	return request, nil
}
