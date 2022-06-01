package models

import "gorm.io/gorm"

//Cuando se convierta en json va a tener ese nombre
type Curso struct {
	//campo anonimo
	gorm.Model
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion,omitempty"`
}
