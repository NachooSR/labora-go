package main
import "fmt"

func main(){
	
	a,b:=10,20

	fmt.Println("Valores iniciales")
	fmt.Println("A= ",a)
	fmt.Println("B= ",b)

    var ptrA * int=&a

	*ptrA,b=b,*ptrA
   

	fmt.Println("Valores intercambiados")
	fmt.Println("A= ",a)

    fmt.Println("B= ",b)
	
}



