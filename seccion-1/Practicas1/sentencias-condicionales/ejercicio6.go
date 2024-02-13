package main

import (
	"fmt"
)

func main() {

	segundos := 0

	fmt.Printf("\nIngrese cantidad de segundos: ")
	fmt.Scanln(&segundos)

	if segundos < 60 {

		fmt.Printf("%d segundos", segundos)

	} else if segundos > 60 && segundos < 3600 {

		minutos := cantMinutos(segundos)
		resto := segundos % 60
		fmt.Printf("\n%d minutos,%d segundos ", minutos, resto)

	} else if segundos >= 3600 && segundos < 86400 {

		horas := cantHoras(segundos)
		restoM := (segundos % 3600) / 60
		restoS := segundos % 60
		fmt.Printf("\n%d horas,%d minutos,%d segundos", horas, restoM, restoS)

	} else if segundos >= 86400 {

		dias := dias(segundos)
		restoH := (segundos % 86400) / 3600
		restoM := (segundos % 3600) / 60
		restoS := (segundos % 60)
		fmt.Printf("\n%d dias,%d horas,%d minutos,%d segundos", dias, restoH, restoM, restoS)

	} else {
		fmt.Printf("\nValor incorrecto")
	}

}

func cantMinutos(segundos int) int {
	return segundos / 60
}

func cantHoras(segundos int) int {
	return segundos / 3600
}
func dias(segundos int) int {
	return segundos / 86400
}
