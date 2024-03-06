package main

import (
	"fmt"
	"math/rand"
	"time"
)

type intSequence interface {
	Next() int
	HolaSoy()
}



type primos struct {
	numeroActual int
}

func (p *primos) Next() int {

	var siguiente int
	for i := p.numeroActual + 1; i <= 50; i++ {
		if isPrimo(i) {
			siguiente = i
			break
		}
	}
	p.numeroActual = siguiente
	return siguiente
}

type creciente struct {
	numeroActual int
}

func (c *creciente) Next() int {
	var siguiente = c.numeroActual
	c.numeroActual++
	return siguiente
}



func isPrimo(numero int) bool {
	if numero <= 1 {
		return false
	}
	for i := 2; i <= numero/2; i++ {
		if numero%i == 0 {
			return false
		}
	}
	return true
}

func (primos) HolaSoy() {
	fmt.Printf("Soy la secuencia de numeros primos\n")
}

func (creciente) HolaSoy() {
	fmt.Printf("Soy la secuencia de numeros crecientes\n")
}

func main() {

	var interfaz intSequence

	numeroRandom := rand.Intn(2)

	if numeroRandom == 0 {
		interfaz = &primos{numeroActual: 1}
	} else {
		interfaz = &creciente{numeroActual: 1}
	}

	interfaz.HolaSoy()

	if (interfaz == &primos{}) {
		for i := 0; i == 47; i++ {
			next := interfaz.Next()
			if next == 0 {
				break
			}
			fmt.Println(next)
		}
	} else {
		for i := 0; i < 50; i++ {
			next := interfaz.Next()
			if next == 0 {
				break
			}
			fmt.Println(next)
		}
	}

	//Tirada de dados
	rand.Seed(time.Now().UnixNano())
	acumulador := 0

	fmt.Printf("\n")
	fmt.Println("Tirada de dados:")
	fmt.Printf("\n")
	for i := 0; i < 30; i++ {
		randomNumber := rand.Intn(6) + 1
		acumulador = acumulador + randomNumber
		fmt.Println(randomNumber)
	}
	fmt.Println("Su tirada de dados acumula un total de: ", acumulador)

}
