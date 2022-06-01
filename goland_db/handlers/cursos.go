package handlers

import (
	"encoding/json"
	"godb/models"
	"godb/repositorios"
	"net/http"
)

func GetCursos(w http.ResponseWriter, r *http.Request) {
	cursos := repositorios.GetCursos()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cursos)
}

func CrearCurso(w http.ResponseWriter, r *http.Request) {
	var curso models.Curso
	json.NewDecoder(r.Body).Decode(&curso)
	//inserta en la base de datos
	curso = repositorios.CrearCurso(curso)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(curso)
}
