package main

import (
	"errors"
	"fmt"


)

func main(){
   numerador,denominador:=0,0

   fmt.Println("Numerador")
   fmt.Scanln(&numerador)
   fmt.Println("Denominador")
   fmt.Scanln(&denominador)

   errorcito,resultado:=division(numerador,denominador)

   if(errorcito!=nil){
	fmt.Println(errorcito)
	return
   }else{
	fmt.Println("El resultado de la division es: ",resultado)
   }

}

func division(numerador,denominador int)(error,float32){
  
  if denominador!=0 {
	var resultado float32
	resultado= float32(numerador)/float32(denominador)
	return nil,resultado
  }else{
	return errors.New("No se puede dividir por 0"),0
  }
}