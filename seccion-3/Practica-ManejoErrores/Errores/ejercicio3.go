package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Hubo un error: ", r)
		}
	}()

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Texto: ")

	texto, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("ERROR")
		return
	}

	texto = strings.TrimSpace(texto)

	numerito := parseo(texto)

	fmt.Println(numerito)
}

func parseo(texto string) int {
	numerito, err := strconv.Atoi(texto)
	if err != nil {
		panic(err)
	}
	return numerito
}
