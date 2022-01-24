package functions

func FuncReturn() {

	n := [] uint8{5,8,6,9,45,255}
	max, min, info := max_min(n) 

	println(max, min, info)
}

func max_min(numbers []uint8) (max, min uint8, info string) {

	info = "esto es un string"

	for _, n := range numbers {
		if n > max {
			max = n
		}

		if n < min {
			min = n
		}
	}

	return
	
}