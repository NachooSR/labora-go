package main

import (
	"fmt"
	"math"
	"math/rand"
)

// Interfaz que pone la condicion para las 3 estructuras
type Condicion interface {
	FullFill(int) bool
}

type secuencia interface {
	Next() int
}

// /Estructura que almacena una interfaz para ser mas flexible
type SecuenciaElegida struct {
	numeroActual int
	condicion    Condicion ///La condicion es una interfaz que luego puede ser la estructura de numeros primos,enteros o pares
}

type numerosPrimos struct{}

type numerosPares struct{}

type numerosEnteros struct{} ///Declaracion de las 3 estructuras

// /Las 3 estructuras con sus metodos fullFill
func (_ numerosPares) FullFill(num int) bool {
	return num%2 == 0
}

func (_ numerosEnteros) FullFill(num int) bool {
	return true
}

func (_ numerosPrimos) FullFill(num int) bool {
	return esPrimo(num)
}
func esPrimo(numero int) bool {
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

/**************************************************************************/

///Estructura y funcion principal

func (s *SecuenciaElegida) Next() int {
	var siguiente int

	for i := s.numeroActual + 1; i < math.MaxInt32; i++ {
		if s.condicion.FullFill(i) { ///si la secuc
			siguiente = i
			break
		}
	}

	s.numeroActual = siguiente

	return siguiente
}

func main() {

	var secuencia SecuenciaElegida

	nroRandom := rand.Intn(3)

	switch nroRandom {
	case 0:
		secuencia = SecuenciaElegida{numeroActual: 0, condicion: numerosPrimos{}}
		break
	case 1:
		secuencia = SecuenciaElegida{numeroActual: 0, condicion: numerosPares{}}
		break
	case 2:
		secuencia = SecuenciaElegida{numeroActual: 0, condicion: numerosEnteros{}}
		break
	}

	for i := 0; i < 50; i++ {
		fmt.Println(secuencia.Next())
	}
}
