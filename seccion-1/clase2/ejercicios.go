package main

import "fmt"

/* EJERCICIOS
1-Escriba un algoritmo para determinar si un número es primo. Recordar que número primo es aquel que es divisible por solo por 1 y por si mismo.
2-Realice un algoritmo que dado un string lo invierta.

*/

func main() {

	//EJERCICIO 1

	var numero int
	fmt.Printf("Ingrese un numero: ")
	fmt.Scanln(&numero)
	check := esPrimo(10)

	if check == false {
		fmt.Printf("Su numero no es primo")
	} else {
		fmt.Printf("Numero primo")
	}

	//EJERCICIO 2

	var palabra string
	fmt.Printf("Ingrese una palabra: ")
	fmt.Scanln(&palabra)

	invertido := stringInvertido([]byte(palabra))

	fmt.Print(string(invertido))

}

func esPrimo(numero int) bool {

	///inicializo un contador
	contador := 0

	if numero <= 1 {
		return true
	}
	//recorre desde 1
	for i := 1; i <= numero; i++ {
		if numero%i == 0 {
			contador++
		}
	}
	if contador > 2 {
		return false
	}

	return true
}

func stringInvertido(palabra []byte) []byte {

	var invertido []byte
	j := len(palabra)
	for i := 0; i < len(palabra); i++ {
		invertido = append(invertido, palabra[j-1])
		j--
	}
	return invertido
}
