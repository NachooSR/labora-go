package main 
import "fmt"

func main(){
	var numero int
	fmt.Printf("Ingrese un numero: ")
    fmt.Scanln(&numero)

	if (numero %2==0 || numero==0) {
		fmt.Printf("Numero par")
	}else{
      fmt.Printf("Numero impar")
	}
}