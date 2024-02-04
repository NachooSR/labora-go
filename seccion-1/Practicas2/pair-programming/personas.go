package main

import (
	"fmt"
	"sort"
	"strings"
)

type Persona struct {
	nombre string
	edad   int
	peso   float32
	altura float32
}

func main() {

	fmt.Printf("\n--------BIENVENIDO, PARA USAR EL PROGRAMA PRIMERO CARGUE 5 PERSONAS------------\n")
	var personas []Persona
	option, option2 := 0, 0

	var aux []Persona

	personas = cargarPersonas(personas)

	for {
		menu()
		fmt.Scanln(&option)
		if option == 4 {
			break
		} else {
			switch option {
			case 1:
				menuBusqueda()
				fmt.Scanln(&option2)
				aux = buscarPersona(option2, personas)
				if(aux!=nil){
					fmt.Println("Esta es la/las personas que coinciden con su busqueda")
					mostrarPersonas(aux)
				}else{
					fmt.Print("\nNo encontramos coincidencias con su busqueda")
				}
				
				break
			case 2:
				mostrarPersonas(personas)
				break
			case 3:
				menuOrdenar()
				fmt.Scanln(&option2)
				ordenarPersonas(option2, personas)
				break
			}

		}
	}

}

func menu() {
	fmt.Println("1- Buscar Personas")
	fmt.Println("2- Ver personas")
	fmt.Println("3- Ordenar personas")
	fmt.Println("4- Salir")
}

func menuBusqueda() {
	fmt.Println("Seleccione como desde buscar las personas")
	fmt.Println("1- Por nombre")
	fmt.Println("2- Por altura")
	fmt.Println("3- Por peso")
	fmt.Println("4- Por edad")
}

func menuOrdenar() {
	fmt.Println("Seleccione como ordenar las personas")
	fmt.Println("1- Por nombre")
	fmt.Println("2- Por altura")
	fmt.Println("3- Por peso")
	fmt.Println("4- Por edad")
}

// Carga de Personas
func cargarPersonas(personas []Persona) []Persona {

	var personita Persona

	///Flag para validar datos de la carga
	flagsito := true

	///variables auxiliares para no cargarlas directamente en la persona
	var alturaPeso float32
	edad := 0

	for i := 0; i < 1; i++ {
		fmt.Print("\nNombre: ")
		fmt.Scanln(&personita.nombre)
		for {
			fmt.Print("\nEdad: ")
			fmt.Scanln(&edad)
			flagsito = validarEdad(edad)
			if flagsito {
				personita.edad = edad
				break
			} else {
				fmt.Println("Dato invalido, por favor pruebe nuevamente")
			}
		}

		for {
			fmt.Print("\nPeso (Kg): ")
			fmt.Scanln(&alturaPeso)
			flagsito = validarAlturaPeso(alturaPeso)
			if flagsito {
				personita.peso = alturaPeso
				break
			} else {
				fmt.Println("Dato invalido, por favor pruebe nuevamente")
			}
		}

		for {
			fmt.Print("\nAltura: ")
			fmt.Scanln(&alturaPeso)
			flagsito = validarAlturaPeso(alturaPeso)
			if flagsito {
				personita.altura = alturaPeso
				break
			} else {
				fmt.Println("Dato invalido, por favor pruebe nuevamente")
			}
		}
        fmt.Println("---------------------")
		personas = append(personas, personita)
	}
	return personas
}

// Muestreo de personas
func mostrarPersonas(personas []Persona) {
	for i := 0; i < len(personas); i++ {
		mostrarPersona(personas[i])
	}
}

// Funcion modularizada
func mostrarPersona(auxiliar Persona) {

	categoria, imc := indiceMasaCorporal(auxiliar.peso, auxiliar.altura)
	fmt.Println("-----------")
	fmt.Println("Nombre: ", auxiliar.nombre)
	fmt.Println("Edad: ", auxiliar.edad)
	fmt.Println("Peso: ", auxiliar.peso)
	fmt.Println("Altura: ", auxiliar.altura)
	switch categoria {
	case 1:
		fmt.Printf("Su indice de masa corporal es: %.2f\n", imc)
		fmt.Printf("Bajo peso: IMC menor a 18.5\n\n")
		break
	case 2:
		fmt.Printf("Su indice de masa corporal es: %.2f\n", imc)
		fmt.Printf("Peso normal: IMC entre 18.5 y 24.9\n\n")
		break
	case 3:
		fmt.Printf("Su indice de masa corporal es: %.2f\n", imc)
		fmt.Printf("Sobrepeso: IMC entre 25 y 29.9\n\n")
		break
	case 4:
		fmt.Printf("Su indice de masa corporal es: %.2f", imc)
		fmt.Printf("Obesidad: IMC mayor a 30\n\n")
		break
	}
}

// *****************ORDENAMIENTOS********************//
func ordenarPersonas(opcion int, personas []Persona) {

	switch opcion {
	case 1:
		personas = ordenarAlfabeticamente(personas)
		mostrarPersonas(personas)
		break
	case 2:
		personas = ordenarPorAltura(personas)
		mostrarPersonas(personas)
		break
	case 3:
		personas = ordenarPorPeso(personas)
		mostrarPersonas(personas)
		break
	case 4:
		personas = ordenarPorEdad(personas)
		mostrarPersonas(personas)
		break
	}
}

func ordenarAlfabeticamente(personas []Persona) []Persona {
	sort.Slice(personas, func(i, j int) bool {
		return personas[i].nombre < personas[j].nombre
	})
	return personas
}

func ordenarPorEdad(personas []Persona) []Persona {
	sort.Slice(personas, func(i, j int) bool {
		return personas[i].edad < personas[j].edad
	})
	return personas
}

func ordenarPorAltura(personas []Persona) []Persona {
	sort.Slice(personas, func(i, j int) bool {
		return personas[i].altura < personas[j].altura
	})
	return personas
}

func ordenarPorPeso(personas []Persona) []Persona {
	sort.Slice(personas, func(i, j int) bool {
		return personas[i].peso < personas[j].peso
	})
	return personas
}

//****************BUSQUEDA DE PERSONAS********************//

func buscarPersona(opcion int, persona []Persona) []Persona {

	/*Las personas pueden compartir nombre,peso, edad y altura
	Por lo que las funciones devuelven un arreglo auxiliar
	Este contendra todas las coincidencias y le mostraremos al usuario la o las personas
	Que comparten el dato ingresado*/

	//*****variables auxiliares*****//
	flagsito:=true
	var auxiliar []Persona
	var nombre string
	var edad int
	var altura float32
	var peso float32
	//*****************************//

	switch opcion {
	case 1:
		fmt.Println("Ingrese el nombre a buscar: ")
		fmt.Scanln(&nombre)
		auxiliar = buscarPersonaNombre(nombre, persona)
		break
	case 2:
		for {
			fmt.Print("\nAltura: ")
			fmt.Scanln(&altura)
			flagsito = validarAlturaPeso(altura)
			if flagsito {
				break
			} else {
				fmt.Println("Dato invalido, por favor pruebe nuevamente")
			}
		}
		auxiliar = buscarPersonaAltura(altura, persona)
		break
	case 3:
		for {
			fmt.Print("\nPeso: ")
			fmt.Scanln(&peso)
			flagsito = validarAlturaPeso(peso)
			if flagsito {
				break
			} else {
				fmt.Println("Dato invalido, por favor pruebe nuevamente")
			}
		}
		auxiliar = buscarPersonaPeso(peso, persona)
		break
	case 4:
		for {
			fmt.Print("\nEdad: ")
			fmt.Scanln(&edad)
			flagsito = validarEdad(edad)
			if flagsito {
				break
			} else {
				fmt.Println("Dato invalido, por favor pruebe nuevamente")
			}
		}
		auxiliar = buscarPersonaEdad(edad, persona)
		break
	}
	return auxiliar
}

func buscarPersonaNombre(nombre string, persona []Persona) []Persona {

	///Arreglo auxiliar
	var auxiliar []Persona

	///En caso de mismo nombre cargamos el arreglo auxiliar con la persona que coincida
	for i := 0; i < len(persona); i++ {
		if strings.EqualFold(persona[i].nombre,nombre) {
			auxiliar=append(auxiliar, persona[i])
		}
	}
	///Lo devolvemos
	return auxiliar
}

func buscarPersonaPeso(peso float32, persona []Persona) []Persona {

	var auxiliar []Persona
	for i := 0; i < len(persona); i++ {
		if persona[i].peso == peso {
			auxiliar=append(auxiliar, persona[i])
		}
	}

	return auxiliar
}
func buscarPersonaAltura(altura float32, persona []Persona) []Persona {

	var auxiliar []Persona

	for i := 0; i < len(persona); i++ {
		if persona[i].altura == altura {
			auxiliar=append(auxiliar, persona[i])
		}
	}

	return auxiliar
}

func buscarPersonaEdad(edad int, persona []Persona) []Persona {

	var auxiliar []Persona
	for i := 0; i < len(persona); i++ {
		if persona[i].edad == edad {
			auxiliar=append(auxiliar, persona[i])
		}
	}

	return auxiliar
}

// //*****************VALIDACIONES*****************////
func validarEdad(edad int) bool {
	flagsito := true
	if edad <= 0 {
		flagsito = false
	}
	return flagsito
}

func validarAlturaPeso(valor float32) bool {
	flagsito := true
	if valor <= 0 {
		flagsito = false
	}
	return flagsito
}

// ////*****CALCULAR IMC************//
func indiceMasaCorporal(peso, altura float32) (int, float32) {
	//1-Bajo peso: IMC menor a 18.5
	//2-Peso normal: IMC entre 18.5 y 24.9
	//3-Sobrepeso: IMC entre 25 y 29.9
	//4-Obesidad: IMC mayor a 30

	categoria := 1

	imc := calculo(peso, altura)

	if (imc >= 18.5) && (imc < 25) {
		categoria = 2
	} else if (imc >= 25) && (imc < 30) {
		categoria = 3
	} else if imc > 30 {
		categoria = 4
	}

	return categoria, imc
}

func calculo(peso, altura float32) float32 {
	return peso / (altura * altura)
}
