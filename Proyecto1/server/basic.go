package main

import (
	"io"
	"log"
	"net/http"
)

//Endpoints

func main() {

	//Handlers
	http.HandleFunc("/usuarios", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			io.WriteString(w, "Servicio de usuarios")
		} else {
			http.NotFound(w, r)
		}

	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/usuarios", 301)
	})

	log.Println("Escuchando en el puerto 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
