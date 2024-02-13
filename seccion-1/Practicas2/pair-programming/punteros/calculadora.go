package main
import "fmt"

func main(){
   a,b:=0,0

   fmt.Println("\nIngrese un numero: ")
   fmt.Scanln(&a)
   fmt.Println("Ingrese un numero: ")
   fmt.Scanln(&b)

   suma,resta,multi,division:=calcular(&a,&b)

   fmt.Println("Suma: ",suma)
   fmt.Println("Resta:  ",resta)
   fmt.Println("Multiplicacion: ",multi)
   fmt.Printf("Division: %.2f\n\n",division)
}

func calcular(a *int,b *int)(int,int,int,float32){
	suma:=  *a + *b
	resta:= *a - *b
    multiplicacion:= *a * *b
	division:= float32(*a)/float32 (*b)
  
    return suma,resta,multiplicacion,division
}