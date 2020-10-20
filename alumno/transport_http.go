//Se llamara a un metodo del service dependiendo de la accion get o post
//LLamadas de los metodos dependiendo de la accion
package alumno

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	kithttp "github.com/go-kit/kit/transport/http"
)

func MakeHttpsHandler(s Service) http.Handler {
	//Creacion de las rutas
	r := chi.NewRouter() //Creando instancia para iniciar el ruteo

	getStudentByIdHandler := kithttp.NewServer(
		makeGetStudentByIdEndPoint(s),
		getStudenByIdRequestDecoder,
		kithttp.EncodeJSONResponse,
	)

	r.Method(
		http.MethodGet, //EL VERBO O METODO GET
		"/{id}",        //El patron o id
		getStudentByIdHandler,
	)
	return r
}

func getStudenByIdRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	studentId, _ := strconv.Atoi(chi.URLParam(r, "id"))
	return getStudentByIDRequest{
		StudentID: studentId,
	}, nil
}
