package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jarelio/tecnicas-de-programacao-ii/backend/controller"
)

func main() {
	router := mux.NewRouter()

	header := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "DELETE", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	controller := &controller.GradesController{}

	router.HandleFunc("/grades", controller.GetGrades).Methods("GET")
	router.HandleFunc("/grades", controller.CreateGrade).Methods("POST")
	router.HandleFunc("/grades/{id:[0-9]+}", controller.GetGrade).Methods("GET")
	router.HandleFunc("/grades/{id:[0-9]+}", controller.DeleteGrade).Methods("DELETE")
	router.HandleFunc("/grades/{id:[0-9]+}", controller.EditGrade).Methods("PUT")
	router.HandleFunc("/grades/student/{student:[a-zA-Z0-9_-]+}", controller.GetGradesByStudent).Methods("GET")

	port := 8000
	log.Printf("Starting grades backend @ %v", port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), handlers.CORS(header, methods, origins)(router)))
}
