package variable

import "fmt"

func Vars() {
	nombre, apellido := "Jose Breiner", "Pai Gonzáles"
	nombre, edad := "Juan", 20

	fmt.Printf("%s %s %d",nombre,apellido, edad)
	
}