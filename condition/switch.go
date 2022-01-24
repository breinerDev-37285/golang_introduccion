package condition

import "fmt"

func SwitchCondition() {


	/*
		No usar break
		fallthrough -> continua evaluando el siguiente case
	
	const a = 6

	switch (a) {
		case 1:
			fallthrough
		case 2:
			fallthrough 
		case 3:
			fallthrough
		case 4:
			fallthrough
		case 5:
			fmt.Println("Estas entre semana")
		case 6:
			fallthrough 
		case 7 :
			fmt.Println("Estas en fin de semana")
		default:
			fmt.Println("Ingresa un valor válido")
	}

	*/

	// asignación de variable y validación con rangos 

	switch a:= 7; {
		case a >= 0 && a <= 5:
			fmt.Println("Estas entre semana") 
		case a >= 6  && a <= 7:
			fmt.Println("Estas en fin de semana")
		default:
			fmt.Println("Ingresa un valor valido")
	}
}