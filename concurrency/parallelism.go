package concurrency

import (
	"fmt"
	"runtime"
	"sync"
)

func isCousin(a uint32) bool{
	c :=0 
	var i uint32
	for i=1; i<= a; i++ {
	
		if a % i == 0 {
			c++
		}
	}

	return c == 2
}

func Parallelism() {
		//parallelism 

	/* wait group espera la finalización de una go rutina*/

	runtime.GOMAXPROCS(1)	// cuantos procesadores debe utilizar
	var wg sync.WaitGroup
	numbers := []uint32{12564,894,265,848,6584,65487,3215,7,987988,65454,897896,6213,6542,84654,986535,204587,23452345,456,567567,4564,1232334,453563456,12341234,2345,564,6585678,789,45}

	wg.Add(len(numbers))	// cuantas go rutinas debe ejecutar

	fmt.Println("Comienza el proceso...")
	for _, v:= range numbers {
		go func(a uint32){
			defer wg.Done()
			fmt.Println(a, isCousin(a))
		}(v)
	} 

	wg.Wait()   // ejecuta el código siguiente si y solo si se termine la ejecución de done
	fmt.Println("Termino el proceso")
}