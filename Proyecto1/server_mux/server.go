package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"go_example/server_mux/handlers"
)

const PORT = ":8000"

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/usuarios", handlers.ObtenerUsuarios).Methods("GET")

	router.HandleFunc("/usuarios", handlers.CrearUsuario).Methods("POST")

	log.Println("Escuchando en el peurto 8000")
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(PORT, nil))
}
