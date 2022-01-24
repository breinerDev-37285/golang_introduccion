package arrays

import "fmt"

func Arrays() {
	var nombres [3]string = [3]string{
		"Juan", 
		"Andres", 
		"Alejando",
	}

	fmt.Printf("%s, size = %d\n",nombres, len(nombres))
	fmt.Printf("%s\n",nombres[0:2]) 	//rango de elementos

}