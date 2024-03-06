package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(1)
	go sumarImpares()
	go sumarPares()
	wg.Wait()
}

func sumarPares() {
	acumulador := 0
	for i := 1; i < 100; i++ {
		if isEven(i) {
			acumulador += i
		}
	}
	fmt.Println("Soy la suma de numeros pares y el resultado es: ", acumulador)
	wg.Done()
}

func sumarImpares() {
	acumulador := 0
	for i := 1; i < 100; i++ {
		if !isEven(i) {
			acumulador += i
		}
	}
	fmt.Println("Soy la suma de numeros impares y el resultado es: ", acumulador)
	wg.Done()
}

func isEven(num int) bool {
	return num%2 == 0
}
