package main

import "fmt"

func main() {
	saludar(20,"Breiner", "Adrian", "Juan", "Pedro")	
}

func saludar(edad uint8, nombres ...string){
	for _, value := range nombres {
		fmt.Printf("Hola %s %d\n", value, edad)
	}
}
