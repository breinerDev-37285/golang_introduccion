package arrays

import "fmt"

// persona es una estructura
type Persona struct {
	nombre 		string
	apellido 	string
	edad 		uint8  
	emails 		[]string
}

func ArraysWithStructs() {
	/* 

		array de objetos 


		var persona1 Persona 
		persona1.nombre = "Alejando"
		persona1.apellido = "Páez"
		persona1.edad = 27 

		fmt.Println(persona1)
	*/

	persona1 := Persona {
		nombre: 	"Alejando",
		apellido: 	"Páez",
		edad:	 	27,
		emails: 	[]string{"email1@mail.com", "email2@mail.com"},
	}

	emails_persona2 := []string{"email1@mail.com", "email2@mail.com"}
	persona2 := Persona {
		"Pablo",
		"Smith",
		30,
		emails_persona2,
	}


	personas := []Persona {
		persona1,
		persona2,
	}

	fmt.Println(personas)
}