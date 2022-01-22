package main

import "fmt"

func main() {

	// funciona igual que el ciclo while

	a := 1
	for a > 0 && a <= 10 {
		fmt.Println(a)
		a++
	}
}