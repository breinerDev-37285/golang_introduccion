package concurrency

import (
	"fmt"
	"time"
)

func ShowNumbers(){
	for i:= 0; i <= 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}