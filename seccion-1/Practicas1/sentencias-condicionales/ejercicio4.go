package main

import (
	"fmt"
	"strings"
)

/*Realice un algoritmo que dado un número te diga si es capicua y si es un string que te diga si es palindromo*/

func main() {

	var palabra string
	fmt.Printf("Ingrese una palabra: ")
	fmt.Scanln(&palabra)

	invertido := stringInvertido([]byte(palabra))

	invertidoString := string(invertido)

	if strings.EqualFold(palabra, invertidoString) {
		fmt.Print("Es palindromo")
	} else {
		fmt.Print("No es")
	}

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
