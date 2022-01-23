package main

import "fmt"

func main() {
	saludar("Breiner",28)


	var saludo string = saludo2("Breiner")
	println(saludo)

	sum := suma(254,255)

	println(sum)

	n := [] uint8{5,8,6,9,45,255}
	max, min, info := max_min(n) 

	println(max, min, info)
}

func saludar(nombre string, edad uint8){
	fmt.Printf("Hola %s tienes %d años\n", nombre, edad)
}

func saludo2(nombre string) string {
	return "Hola "+ nombre
}

func suma(a, b uint16) uint16 {
	return a + b
}

func max_min(numbers []uint8) (max, min uint8, info string) {

	info = "esto es un string"

	for _, n := range numbers {
		if n > max {
			max = n
		}

		if n < min {
			min = n
		}
	}

	return
	
}