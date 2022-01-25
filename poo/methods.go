package poo

import "math"

/*
	firma del método
	para asignar una estructura a una funcion,
	se pasa la estructura antes del nombre de la función

	se recomienda firmar las funciones con punteros para
	no asignar de memoria
*/


func (c *Circle) Area() float64 {
	return c.Radio * math.Pi
}

func (c *Rectangle) Area() float64 {
	return c.Base * c.Hight
}


/*
  	simulación de una clase donde todas las estructuras hereden la función area
	para implementar una interfaz no se pueden pasar como argumento un puntero
*/

func SumAllAreas( figs ...IFigure ) float64{
	area_total := 0.0

	for _, v:= range figs {
		area_total += v.Area()
	}

	return area_total
}