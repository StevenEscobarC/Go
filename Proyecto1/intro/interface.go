package main

import "fmt"

type Animal interface {
	Mover()
	Comer(comida string) string
}

type Perro struct {
	nombre string
	patas  int
}

type Serpiente struct {
	nombre  string
	especie string
}

func (perro *Perro) Mover() {
	fmt.Println("Corriendo")
}

func (perro *Perro) Comer(comida string) string {
	fmt.Println("Comiendo")
	return "Comió"
}

func (serpiente *Serpiente) Mover() {
	fmt.Println("Arrastrandose")
}

func (serpiente *Serpiente) Comer(comida string) string {
	fmt.Println("Comiendo Tsss")
	return "Comió la serpiente"
}

func cualquiera(animal Animal) {
	animal.Mover()
	comioLaPana := animal.Comer("chaca")
	fmt.Println(comioLaPana)
}

func main() {
	perro := Perro{"Zeus", 6}
	cualquiera(&perro)
	serpiente := Serpiente{"ABC", "jajasj"}
	cualquiera(&serpiente)
}
