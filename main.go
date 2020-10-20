package main

import (
	"net/http"

	"github.com/RestApiELearningLP2/alumno"
	"github.com/RestApiELearningLP2/database"
	"github.com/go-chi/chi"
)

func main() {
	databaseConnection := database.InitDB()

	defer databaseConnection.Close()

	var studentRepository = alumno.NewRepository(databaseConnection)
	var studentService alumno.Service
	studentService = alumno.NewService(studentRepository)

	r := chi.NewRouter()

	r.Mount("/alumnos", alumno.MakeHttpsHandler(studentService))
	http.ListenAndServe(":3000", r)
}
