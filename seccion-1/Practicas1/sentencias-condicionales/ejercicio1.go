package main
import "fmt"

func main(){
	var numero int
	fmt.Printf("Ingrese un numero: ")
    fmt.Scanln(&numero)

	if(numero==0){
    fmt.Print("Numero es igual 0")
	}else if(numero>0){
		fmt.Print("Numero mayor a 0")
	}else{
		fmt.Print("Numero menor a 0")
	}
}