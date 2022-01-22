package main

import (
	"fmt"
	"strconv"
)

func main() {

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

}