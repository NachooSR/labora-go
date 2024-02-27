package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/scanner"
)

func main(){
       

	scanner := bufio.NewScanner(os.Stdin)
    fmt.Printf("\nPor favor, ingresa un texto:\n")
    scanner.Scan()
    texto := scanner.Text()
	 
	mapita:=contadorPalabras(texto)
	 fmt.Printf("\n\n")
	 fmt.Println(mapita)

}

func contadorPalabras(texto string) map[string]int{
    
	cantidad:=strings.Fields(texto)
    
    mapa:=make(map[string]int)

	for _, palabra:=range cantidad{
      mapa[palabra]++
	}
	return mapa

}