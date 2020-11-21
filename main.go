package main

import (
	"encoding/json"
	"net/http"

	"github.com/RestApiELearningLP2/database"
	"github.com/RestApiELearningLP2/middlew"
	"github.com/RestApiELearningLP2/usuario"
	"github.com/go-chi/chi"
)

func main() {
	databaseConnection := database.InitDB()

	defer databaseConnection.Close()

	var studentRepository = usuario.NewRepository(databaseConnection)
	var studentService usuario.Service
	studentService = usuario.NewService(studentRepository)

	r := chi.NewRouter()

	r.Mount("/usuario", middlew.ValidoJWT(usuario.MakeHttpsHandler(studentService)))
	http.ListenAndServe(":3000", r)
}

func mostrarMensaje(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusCreated, map[string]string{"mensaje": "Esta es"})
}

func respondWithJSON(w http.ResponseWriter, cod int, playload interface{}) {
	response, _ := json.Marshal(playload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(cod)
	w.Write(response)
}
