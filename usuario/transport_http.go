//Se llamara a un metodo del service dependiendo de la accion get o post
//LLamadas de los metodos dependiendo de la accion
package usuario

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
)

func MakeHttpsHandler(s Service) http.Handler {
	//Creacion de las rutas
	r := chi.NewRouter() //Creando instancia para iniciar el ruteo

	getStudentByIdHandler := kithttp.NewServer(
		registerUserEndPoint(s),
		registerUserRequestDecoder,
		kithttp.EncodeJSONResponse,
	)

	r.Method(http.MethodPost, "/", getStudentByIdHandler)
	return r
}

func registerUserRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := registerUserRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	return request, err
}
