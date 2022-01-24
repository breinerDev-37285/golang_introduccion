package functions

import "fmt"

func SpreedFunc() {
	saludar(20,"Breiner", "Adrian", "Juan", "Pedro")	
}

func saludar(edad uint8, nombres ...string){
	for _, value := range nombres {
		fmt.Printf("Hola %s %d\n", value, edad)
	}
}
