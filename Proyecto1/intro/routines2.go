package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	//Simular consulta
	//Rutina, trabajan en paralelo
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second * 2)
			fmt.Println("Tarea 1")
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second * 2)
			fmt.Println("Tarea 2")
		}
		wg.Done()
	}()

	fmt.Println("Esperando las rutinas")
	wg.Wait()
	fmt.Println("Terminaron las rutinas")
	fmt.Println("Termino el main")
}
