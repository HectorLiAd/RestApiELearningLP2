package main

import (
	"net/http"

	"github.com/RestApiELearningLP2/database"
	"github.com/RestApiELearningLP2/middlew"
	"github.com/RestApiELearningLP2/usuario"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	databaseConnection := database.InitDB()

	defer databaseConnection.Close()

	var studentRepository = usuario.NewRepository(databaseConnection)
	var studentService usuario.Service
	studentService = usuario.NewService(studentRepository)

	r := chi.NewRouter()

	r.Mount("/usuariosistema", middlew.ValidoJWT(usuario.MakeHTTPSHandler(studentService)))
	http.ListenAndServe(":3000", r)
}
