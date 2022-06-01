package main

import (
	"fmt"
	"time"
)

func main() {
	//Simular consulta
	//Rutina, trabajan en paralelo
	go tarea1()
	go tarea2()

	fmt.Println("Termino el ciclo")
	fmt.Scanln()
}

func tarea1() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 2)
		fmt.Println("Tarea 1")
	}
}

func tarea2() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 2)
		fmt.Println("Tarea 2")
	}
}
