package main 
import "fmt"

/*
# Descripción del problema

Escribe un programa que calcule el área y el perímetro de un rectángulo. El programa debe pedir al usuario que introduzca la longitud y la anchura del rectángulo y luego calcular y mostrar el área y el perímetro.

## Especificaciones:

1. Solicita al usuario que ingrese la longitud y la anchura del rectángulo.
2. Calcula el área del rectángulo (longitud * anchura).
3. Calcula el perímetro del rectángulo (2 * (longitud + anchura)).
4. Muestra los resultados (área y perímetro) al usuario.

*/

func main(){

	var num1,num2 float32

	fmt.Printf("\nLongitud del rectangulo: ")
    fmt.Scanln(&num1)
	fmt.Printf("\nArea: ")
	fmt.Scanln(&num2)
    
	num1=areaR(num1,num2)
	num2=perímetroR(num1,num2)

	fmt.Printf("\nArea del rectangulo %.2f",num1)
	fmt.Printf("\nPerimetro del rectangulo %.2f",num2)
}

func areaR(longitud,area float32)float32{
	return longitud*area
}
func perímetroR(longitud,area float32)float32{
	return (2*(longitud+area))
}