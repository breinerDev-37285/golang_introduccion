package panics

import "fmt"

func recovers(){

	/*recupera la información enviada por panic*/

	if r := recover(); r != nil {
		fmt.Printf("%T -> %d", r, r)
	}
}


func F(i uint8){
	
	defer recovers()

	fmt.Println("llamando a g() ")
	g(i)

	fmt.Println("cool, se recupero")
}

func g(i uint8) {

	if i > 3 {
		fmt.Println("Aaaaahhh!")
		panic(i)
	}
	
	fmt.Println("valor de i=", i)
}