package main

import (
	"fmt"

	"math"
	"math/rand"

	"golang.org/x/tools/go/analysis/passes/nilfunc"
)

type Audio struct {
	duracion int
}

type Imagen struct {
	size string
}

type Texto struct {
	palabras int
}

type Documento interface {
	Medidas()
}

func (t *Texto) Medidas() {
	fmt.Println("Documento de tipo texto con:", t.palabras, "palabras")

}

func (x Imagen) Medidas() {
	fmt.Println("Imagen con resolucion: ", x.size)
}

func (x Audio) Medidas() {
	fmt.Println("Archivo audio:", x.duracion, "segundos")
}

func main() {

	option := 0

	nombre := ""
	documentos := make(map[string]Documento)

	for option != 4 {
		fmt.Printf("\n\n")
		fmt.Println("Seleccione tipo de documento: ")
		menu()
		fmt.Scanln(&option)

		switch option {
		case 1:
			fmt.Printf("\n\n")
			fmt.Println("Ingrese nombre del documento")
			fmt.Scanln(&nombre)
			nombreExtendido := nombre + ".txt"
			palabras := rand.Intn(math.MaxInt16)
			documentos[nombreExtendido] = &Texto{palabras: palabras}

			break
		case 2:
			fmt.Printf("\n\n")
			fmt.Println("Ingrese nombre del documento")
			fmt.Scanln(&nombre)
			nombreExtendido := nombre + ".mp3"
			duracion := rand.Intn(59)
			documentos[nombreExtendido] = &Audio{duracion: duracion}
			break
		case 3:
			var medida string
			fmt.Printf("\n\n")
			fmt.Println("Ingrese nombre del documento")
			fmt.Scanln(&nombre)
			nombreExtendido := nombre + ".jpg"
			fmt.Println("Ingrese medida de su imagen (Ej: 1920x1080): ")

			fmt.Scanln(&medida)
			documentos[nombreExtendido] = &Imagen{size: medida}
			break
		case 4:
			break
		default:
			fmt.Println("Opcion incorrecta")
			nombre = ""
		}
	}

	fmt.Printf("\n\n\nDocumentos agregados\n")
	fmt.Printf("-------------------\n")
	for nombre, doc := range documentos {

		fmt.Println("Doc name: ", nombre)
		doc.Medidas()
		fmt.Printf("\n------------\n")
	}

}

func menu() {
	fmt.Println("1-Texto")
	fmt.Println("2-Audio")
	fmt.Println("3-Imagen")
	fmt.Println("4-Salir")
}
