package main

import (
	"introduccion/panics"
)

/*
func main() {
	f1 := poo.Circle{Radio: 6}
	f2 := poo.Rectangle{Base: 2, Hight: 10}

	areas := poo.SumAllAreas(&f1, &f2)

	fmt.Printf("El área de la figura 1 es %.2f\n",f1.Area())
	fmt.Printf("El área de la figura 2 es %.2f\n",f2.Area())
	fmt.Printf("El área total es %.2f\n",areas)

}
*/

/*
func main(){


	// todos los diferidos se ejecutan al final
	// el primer defer declarado sera el ultimo en ejecutarse
	// los diferidos se ejecutan aunque la aplicación entre en pánico


	//conecta y desconecta la base de datos
	defers.ConnectDB()					//1
	defer defers.DisconnectDB()			//4

	// consulta la información
	defers.ConsultData()				//2
	defer defers.CloseConsultData()		//3

}
*/

func main() {
	// panics.Division(4,0)
	panics.F(4)
}