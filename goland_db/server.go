package main

import (
	"godb/db"
	"godb/handlers"
	"godb/middlewares"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const PORT = ":8000"

func main() {
	db.InitDB()

	router := mux.NewRouter()

	router.HandleFunc("/", middlewares.Middleware(handlers.GetCursos)).Methods("GET")
	router.HandleFunc("/", middlewares.Middleware(handlers.CrearCurso)).Methods("POST")
	http.Handle("/", router)

	log.Println("Escuchando en el puerto ", PORT)
	log.Fatal(http.ListenAndServe(PORT, nil))
}
