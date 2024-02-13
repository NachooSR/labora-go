package main
import "fmt"

/*Dados tres valores ambientales de temperatura, desarrollar un algoritmo para calcular e informar la suma y el promedio de dichos valores.*/

func main(){

	acumulador,temp:=0,0
	


    for i := 0; i < 3; i++ {
	   fmt.Print("Ingrese un numero: ")
	   fmt.Scanln(&temp)
	   acumulador=acumulador+temp
	   fmt.Print("\n")	
	}

	fmt.Print("La suma de las temperaturas es: ",acumulador)
	fmt.Printf("\nEl promedio es: %.2f",promedio(acumulador))

}


func promedio(suma int) float64{

 return float64(suma)/3
}