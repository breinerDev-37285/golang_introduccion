package variable

import "fmt"

func Formats() {
	const a int = 10345345345
	const b int8 = 34
	const isTall bool = false

	// casting

	c := a + int(b)
	fmt.Printf("%d\t%t\n", c, isTall)
	fmt.Printf("tipo: %T\t%T\n", isTall,c)

	/* 
		prioridad aritmética 
		* / + -
	*/

	d := 6 + 6 * 6 - 6
	fmt.Printf("%d\n",d)

	// módulos  
	e := 6 % 2 
	fmt.Println(e)

}
