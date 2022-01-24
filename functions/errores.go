package functions

import (
	"errors"
	"fmt"
)


func ErrFunc() {
	resultado,err := division(100,0) 
	
	if err == nil {
		fmt.Println("Error",err)
	}

	fmt.Printf("%.2f\n", resultado)
}


func division (dividendo, divisor float64) (resultado float64, err error) {
	
	if(dividendo == 0) {
		err = errors.New("Divisor no puede ser cero")
	}else {
		resultado = dividendo / divisor
	}
	return 
}