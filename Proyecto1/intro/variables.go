package main

import "fmt"

func main() {

	// //numericos int float32 float64
	// var numero int //declarar
	// numero = 10    //inicializar
	// fmt.Println(numero)
	// flotante := 10.1 //declarar e inicializar solo cuando se esta creando una variable nueva
	// fmt.Println(flotante)

	// //cadenas string
	// cadena := "hola mundo"
	// fmt.Println(cadena)

	// //booleanos bool
	// bandera := true
	// fmt.Println(bandera)

	// suma := numero + int(flotante)
	// fmt.Println(suma)

	// //arreglos - slices
	// var arreglo [5]int
	// fmt.Println(arreglo)

	//letras := [2]string{"a", "b"}
	// fmt.Println(letras)

	// //slices
	// slice := make([]int, 5, 100)
	// fmt.Println(slice)
	// fmt.Println(len(slice))
	// fmt.Println(cap(slice))

	// slice2 := []int{1, 2, 3}
	// copy(slice2, slice)
	// fmt.Println(slice2)

	//mapas - maps
	mapa := make(map[string]int)
	mapa["uno"] = 1
	mapa["dos"] = 2
	mapa["xos"] = 70
	fmt.Println(mapa)
	fmt.Println(mapa["dos"])

	// fmt.Println("--------------------------------------")
	// for i, v := range letras {
	// 	fmt.Printf("Indice: %d - valor: %s ", i, v)
	// }

	fmt.Println("--------------------------------------")
	for k, v := range mapa {
		fmt.Printf("%s -> %d \n", k, v)
	}

	var nombre string
	fmt.Println("Inserte el nombre")
	fmt.Scanf("%s", &nombre)

	fmt.Println("su nombre es " + nombre)

	var edad int
	fmt.Println("Inserte la edad")
	fmt.Scanf("%s", &edad)

	fmt.Println("su edad es ", edad)

	//punteros *

	//structs

}
