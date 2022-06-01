package repositorios

import (
	"godb/db"
	"godb/models"
	"log"
)

func GetCursos() (cursos []models.Curso) {
	db := db.GetDB()
	db.Find(&cursos)
	return
}

func CrearCurso(body models.Curso) (curso models.Curso) {
	db := db.GetDB()
	db.Create(&body)
	log.Println(body)
	//db.Last(&curso)
	curso = body

	return
}
