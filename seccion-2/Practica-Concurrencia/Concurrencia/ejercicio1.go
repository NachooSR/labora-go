package main

import (
	"fmt"
	"time"
)

var channel chan int = make(chan int)

func main() {

	arreglo := []int{3,4,6,1,5,2,7,8,9,10}
	sleepSort(arreglo)
	/*****************NO LO PUDE RESOLVER TODAVIA***********************/

}

func sleepSort(arreglo []int) {

	///Declaramos un canal

	///Un for que itere sobre estos valores
	for _, i := range arreglo {
		go func(num int) { ///Una goroutine con los valores que vamos recibiendo del arreglo

			///Por cada valor del arreglo se abre una goRoutine
			///A mayor numero mayor es el tiempo que permanece dormido ese hilo, por eso mismo se ordenan
			time.Sleep(time.Duration(num) * time.Millisecond)
			channel <- num

		}(i)
	}

	for i := 0; i < len(arreglo); i++ {
		fmt.Print(<-channel, " ")
	}
	fmt.Println()
}
