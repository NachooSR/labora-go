package main

import "fmt"

/*Expresar a un número cualquiera como la suma de una serie de números siguiendo las restricciones impuestas a continuación.
RESTRICCIONES
S1 ≤ 50
S2 ≤ 50
S3 ≤ 600         TODOS MAYORES A 0
S4 ≤ 800
S5 < (Infinito)
*/

func main() {
     
	numero:=0 
	 for {
		fmt.Printf("\nIngrese un numero: ")
		fmt.Scanln(&numero)
		if(numero>0){
          break
		}else{
			fmt.Print("\nDato invalido, por favor ingrese nuevamente")
		}
	 }
	 descomponer(numero)
     
}
func descomponer(numero int){
	
	resto:=0
	if(numero<=50){
		fmt.Printf("\n %d = %d + 0 + 0 + 0 + 0",numero,numero)
	}else if(numero>50 && numero<=100){
		
		resto=numero-50
		fmt.Printf("\n%d= 50 + %d + 0 + 0 + 0",numero,resto)
	
	}else if(numero>100 && numero<=700){
        resto=numero-100
		fmt.Printf("\n%d= 50 + 50 + %d + 0 + 0",numero,resto)
	}else if(numero>700 && numero<=1500){
		resto=numero-700
		fmt.Printf("\n%d= 50 + 50 + 600 + %d + 0",numero,resto)
	}else{
		resto=numero-1500
		fmt.Printf("\n%d= 50 + 50 + 600 + 800 + %d",numero,resto)
	}
}
