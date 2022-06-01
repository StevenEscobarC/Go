package handlers

import (
	"encoding/json"
	"go_example/server_mux/models"
	"net/http"
)

func ObtenerUsuarios(w http.ResponseWriter, r *http.Request) {
	usuarios := make([]string, 0)
	usuarios = append(usuarios, "Juan")
	usuarios = append(usuarios, "Pedro")
	usuarios = append(usuarios, "Maria")

	json.NewEncoder(w).Encode(usuarios)
}

func CrearUsuario(w http.ResponseWriter, r *http.Request) {
	var usuario models.Usuario
	err := json.NewDecoder(r.Body).Decode(&usuario)

	if err != nil {
		//w.WriteHeader(400)
		http.Error(w, "Error al leer el usuario", http.StatusBadRequest)
		return
	}

	//Aqui se escribiría a la base de datos

	json.NewEncoder(w).Encode(usuario)
}
