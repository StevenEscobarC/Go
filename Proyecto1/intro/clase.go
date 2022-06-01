package main

import "fmt"

type Persona struct {
	nombre string
	edad   int
	// datos    []string
	// arreglo  [10]string
	// mapa     map[string]string
	// mascotas []Mascotas
}

type Mascotas struct {
	nombre string
	edad   int
}

//funcion que pertenece a la estructura persona
func (persona Persona) saludar() {
	fmt.Println("Hola " + persona.nombre)
}

//pasar por referencia, puntero. sin asterisco es por valor(hace copia y no cambia la referencia en memoria)
func (persona *Persona) incrementarEdad(cant int) {
	persona.edad += cant
}

type Empleado struct {
	//herencia
	Persona
	cargo string
}

func main() {
	persona := Persona{nombre: "Juan", edad: 30}
	fmt.Println(persona)
	persona.nombre = "otro"
	fmt.Println(persona)
	persona.incrementarEdad(10)
	persona.saludar()
	fmt.Println(persona)

	empleado := Empleado{}
	empleado.cargo = "Programador"
	empleado.nombre = "Steven"
	empleado.edad = 30

	empleado.incrementarEdad(10)

	fmt.Println("----------------")
	fmt.Println(empleado)
}
