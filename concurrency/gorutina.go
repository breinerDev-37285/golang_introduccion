package concurrency

import (
	"fmt"
	"time"
)

func showNumbers(){
	for i:= 0; i <= 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}


func Concurrency(){
	// concurrency 
	var h string
	go showNumbers()

	fmt.Print("Digita algo: ")
	fmt.Scan(&h)
	fmt.Println("Digistaste ", h) 
	
}