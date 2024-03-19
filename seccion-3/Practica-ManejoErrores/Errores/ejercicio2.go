package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main(){
	
   reader:= bufio.NewReader(os.Stdin)
   fmt.Printf("Texto: ")
   
   ///Convertir el reader en string
   texto,err:= reader.ReadString('\n')
    if(err!=nil){
	fmt.Println("ERROR")
	return
    }
    
   ///Eliminamos los espacios vacios --> DEL PRINCIPIO Y DEL FINAL
   texto=strings.TrimSpace(texto)

   errorcito:=verificacion(texto)

   if(errorcito!=nil ){
         fmt.Println(errorcito)
    return   
   }
   fmt.Println("Programa exitoso sin errores")
}

func contarLetras(texto string)int{
	contador:=0
	for _, letra:= range texto{
        if unicode.IsLetter(letra) {
			contador++
		}
	}
	return contador
}

func verificacion(texto string)(error){
    if(texto==""){
	   return errors.New("ERROR-Texto vacio")
	}else if(contarLetras(texto)>100){
		return errors.New("ERROR-El texto excede el limite de caracteres")
	}
	return nil
}