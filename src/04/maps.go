package main

import "fmt"

func main() {
	// make(tipo de dato)
	// idiomas[es] = "español"

	/* 
		forma 1

		idiomas := make(map[string]string)

		idiomas["es"] = "español"
		idiomas["en"] = "ingles"
		idiomas["fr"] = "francés"

		fmt.Println(idiomas["es"]) 
	*/


	idiomas := map[string]string {
		"es": "español",
		"en": "ingles",
		"fr": "francés",
	}

	fmt.Println(idiomas)

	/* 
		delete(idiomas,"en")
		fmt.Println(idiomas)
	*/
		
	// comprobar existencia de clave 
	if idioma, ok := idiomas["pt"]; ok {
		fmt.Printf("la posición %s si existe", idioma)
	}else {
		fmt.Printf("La posición %s no existe", idioma)
	}

}