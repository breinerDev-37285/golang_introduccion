package concurrency

import "fmt"

/*
	todas las go rutinas realizan trabajos independientes, sin embargo, en algunas ocasiones es necesario que
	estas se comuniquen, para realizar esto se usan los canales
*/

func Channels(){
	/* simulación para una cola de fotocopiado */
	
	// libros 
	bodegaOrigen := []string{"php", "html", "css", "java", "sql"}
	bodegaDest := []string{}

	// se crea un canal -> make (chan tipo de dato)
	fotocopiadora := make(chan string)
	fotocopiado := make(chan string)  
	

	// se envía el libro a la fotocopiadora
	go func(){
		for _, libro := range bodegaOrigen {
			fotocopiadora <- libro
		}
	}()
	
	// se deja en la bodega de destino
	go func(){
		var libro string
		
		for {
			libro = <-fotocopiadora // se entrega el contenido del canal a la variable
			fotocopiado <- libro	 // se entrega el contenido del libro al canal fotocopiado

			bodegaDest = append(bodegaDest, libro) // se agrega el libro a la bodega de destino
			
			if len(bodegaOrigen) == len(bodegaDest){	
				close(fotocopiado) // se cierra el canal
			}
		}
	}()


	fmt.Println("** Listado de libros fotocopiados **")
	for {
		if libro, ok:= <-fotocopiado; ok {
			fmt.Println(libro)
		}else {
			break
		}
	}
}