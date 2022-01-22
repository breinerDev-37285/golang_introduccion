package main

import "fmt"

func main () {
	
	/*
		arrays de tamaño variable  - listas
		principios 

		1. Apuntador a un array 
		2. Tamaño 
		3. Capacidad 

	*/

	// forma 1 
	// slice = append(slice, valor)
	var nombres []string 

	nombres = append(nombres, "Andres")
	nombres = append(nombres, "Carlos")
	nombres = append(nombres, "Adrian")
	nombres = append(nombres, "Daniel")
	nombres = append(nombres, "Michelle")
	nombres = append(nombres, "Michael")
	nombres = append(nombres, "Breiner")


	fmt.Printf("%s, size=%d, capacity=%d\n",nombres,len(nombres), cap(nombres))


	// forma 2 
	// make(tipo de dato, tamaño inicial, capacidad inicial)
	apellidos := make([]string, 0)

	fmt.Println(apellidos)


	
	// forma 3 
	// declaración inmediata 

	days := []string {
		"lunes", "martes","miércoles", "jueves", "viernes",
	}

	fmt.Printf("%s	size=%d	capacity=%d", days, len(days), cap(days))
}