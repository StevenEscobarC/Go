package db

import (
	"fmt"
	"godb/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	user     = "root"
	password = "admin123"
	dbname   = "cursos_db"
	host     = "127.0.0.1"
	port     = "3306"
)

var db *gorm.DB

func Conectar() *gorm.DB {
	dbCon := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)
	db, err := gorm.Open(mysql.Open(dbCon), &gorm.Config{})

	if err != nil {
		log.Println("Error al conectar a la base de datos")
		panic(err)
	}
	log.Println("Conectado a la base de datos")

	return db
}

//para mapear las clases en tablas en la base de datos
func Migrar() {
	db.AutoMigrate(
		//& referencia
		&models.Curso{},
	)
}

func InitDB() {
	db = Conectar()
	Migrar()
}

func GetDB() *gorm.DB {
	return db
}
