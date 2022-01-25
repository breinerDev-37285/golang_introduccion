package panics

import "fmt"

func Division(dividendo,divisor int) {

	/* Los diferidos se ejecutan siempre y cuando se declaren antes del panic*/

	defer fmt.Println("esto se ejecutara pase lo que pase")
	
	if divisor == 0 {
		panic("No se puede dividir para cero")
	}
	fmt.Println(dividendo, divisor)
}