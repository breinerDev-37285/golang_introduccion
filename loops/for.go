package loops

import (
	"fmt"
	"strconv"
)

func ForClassic() {

	/* 	
		var c string = ""

		for i := 0; i < 20; i++ {
			if i % 2 == 0 {
				c += strconv.Itoa(i)+"\n"
			}

			if i == 9 {
				break
			}
		}

		fmt.Printf("%s",c)

	*/

	números := [3][2]int{}
	valor := 0


	for i := 0; i < len(números); i++ {
		for j := 0; j < len(números[i]); j++ {
			valor++ 
			números[i][j] = valor
		}
	}

	c := ""

	for i := 0; i < len(números); i++ {
		c+="[ "
		for j := 0; j < len(números[i]); j++ {
			c += strconv.Itoa(números[i][j])+" "
		}
		c += "]\n"
	}

	fmt.Printf("%s", c)
}