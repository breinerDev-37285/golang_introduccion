package main

import (
	"fmt"
	"introduccion/poo"
)


func main() {
	/*
		poo
	*/

	f1 := poo.Circle{Radio: 6}
	f2 := poo.Rectangle{Base: 2, Hight: 10}

	areas := poo.SumAllAreas(&f1, &f2)

	fmt.Printf("El área de la figura 1 es %.2f\n",f1.Area())
	fmt.Printf("El área de la figura 2 es %.2f\n",f2.Area())
	fmt.Printf("El área total es %.2f\n",areas)
	
}


