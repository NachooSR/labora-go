package main

import "fmt"

/*
Realice un algoritmo para determinar si dos números son amigos. Un número es amigo de otro cuando la suma de sus divisores propios es igual al otro número.
Sean X1, X2, XN todos divisores propios de X
Sean Y1, Y2, YN todos divisores propios de Y
Si X e Y son amigos entonces X1 + x2 +…+ XN es igual a Y e Y1 + Y2 +…+ XN es igual a X
*/

func main() {

	numero1, numero2 := 0, 0
	fmt.Printf("\nIngrese un numero: ")
	fmt.Scanln(&numero1)
	fmt.Printf("\nIngrese otro numero: ")
	fmt.Scanln(&numero2)

	total1 := numerosDivisores(numero1)
	total2 := numerosDivisores(numero2)

	if total1 == numero2 && total2 == numero1 {
		fmt.Printf("\nNumeros amigos")
	}

}

func numerosDivisores(numero int) int {
	total := 0
	var arregloDivisores []int

	//Generamos un arreglo con los divisores del numero
	for i := 1; i < numero; i++ {
		if numero%i == 0 {

			arregloDivisores = append(arregloDivisores, i)
		}
	}

	//Recorremos los valores que guardamos anteriormente y los acumulamos
	for i := 0; i < len(arregloDivisores); i++ {
		total = total + arregloDivisores[i]
	}

	return total
}
