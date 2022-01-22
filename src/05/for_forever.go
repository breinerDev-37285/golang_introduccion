package main

func main() {

	// se ejecuta indefinidamente, equivalente a while(true)

	c := 0

	for {
		println(c)
		c++

		if c == 10 {
			break
		}
	}
	
}