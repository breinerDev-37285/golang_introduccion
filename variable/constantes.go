package variable

import "fmt"

func Constants() {
	const pi float32 = 3.141592
	const edad uint8 = 26

	fmt.Printf("El valor de pi es = %.2f -> la edad %d",pi, edad)
}
