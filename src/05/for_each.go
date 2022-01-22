package main


func main() {

	personas := []string{"Andres", "Michael", "Jordan", "Carl"}


	// for indice, valor := range array
	
	for indice, valor := range personas {
		println(indice, valor)
	}

	println("\n")

	// acceder solo a los valores, usar underscore
	for _, valor := range personas {
		println(valor)
	}

	println("\n")

	//acceder solo al indice, omitir segundo parámetro

	for i := range personas {
		println(i)
	}

	println("\n")

	// recorrer caractares ASCII

	frase := "Hola, mundo"
	for i, v := range frase {
		println(i,string(v), v)
	}

}