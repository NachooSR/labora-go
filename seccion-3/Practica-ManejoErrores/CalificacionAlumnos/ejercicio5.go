package main

import (
	"errors"
	"fmt"
	"strconv"
)

type Alumno struct {
	id                  int
	legajo              string
	nombre              string
	apellido            string
	materiaCalificacion map[string]float32
}

// /Cada vez que se crea un alumno se incrementa esta variable
// /Y es alojada en el Id del alumno
var idAutoincremental int

// /Cuando nuestro programa recien inicia esta variable sera falsa
// /Para evitar errores como ver un arreglo vacio u operaciones con el mismo
var alumnosCargados bool = true

///Esta puesta en true porque cargue 5 alumnos a mano para ya probar las funcionalidades

/*
arregloAlumnos := []Alumno{
        {
            id:       1,
            legajo:   "001",
            nombre:   "Juan",
            apellido: "Perez",
            materiaCalificacion: map[string]float32{
                "Matemáticas":  8.5,
                "Ciencias":     7.2,
                "Literatura":   9.0,
            },
        },
        {
            id:       2,
            legajo:   "002",
            nombre:   "María",
            apellido: "Gomez",
            materiaCalificacion: map[string]float32{
                "Matemáticas":  7.8,
                "Ciencias":     8.0,
                "Literatura":   8.5,
            },
        },
        {
            id:       3,
            legajo:   "003",
            nombre:   "Pedro",
            apellido: "Rodríguez",
            materiaCalificacion: map[string]float32{
                "Matemáticas":  6.5,
                "Ciencias":     7.0,
                "Literatura":   6.8,
            },
        },
        {
            id:       4,
            legajo:   "004",
            nombre:   "Ana",
            apellido: "Martínez",
            materiaCalificacion: map[string]float32{
                "Matemáticas":  9.0,
                "Ciencias":     8.5,
                "Literatura":   8.2,
            },
        },
        {
            id:       5,
            legajo:   "005",
            nombre:   "Carlos",
            apellido: "López",
            materiaCalificacion: map[string]float32{
                "Matemáticas":  7.0,
                "Ciencias":     6.5,
                "Literatura":   7.8,
            },
        },
    }

*/

func main() {

	//var alumnos []Alumno
	var option int

	fmt.Printf("\nBienvenido a Labora School\n")
	var arregloAlumnos []Alumno

	arregloAlumnos = []Alumno{
		{
			id:       1,
			legajo:   "001",
			nombre:   "Juan",
			apellido: "Perez",
			materiaCalificacion: map[string]float32{
				"Matemáticas": 8.5,
				"Ciencias":    7.2,
				"Literatura":  9.0,
			},
		},
		{
			id:       2,
			legajo:   "002",
			nombre:   "María",
			apellido: "Gomez",
			materiaCalificacion: map[string]float32{
				"Matemáticas": 7.8,
				"Ciencias":    8.0,
				"Literatura":  8.5,
			},
		},
		{
			id:       3,
			legajo:   "003",
			nombre:   "Pedro",
			apellido: "Rodríguez",
			materiaCalificacion: map[string]float32{
				"Matemáticas": 6.5,
				"Ciencias":    7.0,
				"Literatura":  6.8,
			},
		},
		{
			id:       4,
			legajo:   "004",
			nombre:   "Ana",
			apellido: "Martínez",
			materiaCalificacion: map[string]float32{
				"Matemáticas": 9.0,
				"Ciencias":    8.5,
				"Literatura":  8.2,
			},
		},
		{
			id:       5,
			legajo:   "005",
			nombre:   "Carlos",
			apellido: "López",
			materiaCalificacion: map[string]float32{
				"Matemáticas": 7.0,
				"Ciencias":    6.5,
				"Literatura":  7.8,
			},
		},
	}

	for option != 4 {

		///Mostramos las opciones al usuario
		menuPrincipal()

		var input string
		fmt.Scan(&input)

		///Escaneamos lo que ingresa si no es un numero ERROR
		///Si es un numero lo parseamos
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println(" **Error: Debe ingresar un número**")
			continue
		}

		///Parseo de la opcion ingresada
		option = num

		if (option == 2 || option == 3) && alumnosCargados == false {
			fmt.Println(" ******** ERROR ************")
			fmt.Println(" Debe cargar alumnos\n")
			continue
		}

		switch option {
		case 1:
			arregloAlumnos = CargaAlumnos(arregloAlumnos)
			break

		case 2:
			optionCase2 := 0
			muestraAlumnos(arregloAlumnos)
			///Hasta aca bien

			fmt.Printf("\n Ver promedios--> (0-Si 1-No): ")
			fmt.Scan(&optionCase2)
			if optionCase2 == 0 {
				promedios(arregloAlumnos)
			}

			break

		case 3:
			////Variables auxiliares
			var legajo string
			var auxiliar Alumno
			var err error
			var banderin bool = false

			var option2 int
			///////////////////////////////
			for banderin == false {

				option2 = 0

				for option2 == 0 {
					fmt.Printf("\n Ingrese legajo: ")
					fmt.Scan(&legajo)
					err, auxiliar = buscarAlumno(legajo, arregloAlumnos)

					///Si no existe el Legajo, preguntamos si quiere probar otro
					if err != nil {
						fmt.Println(err)
						fmt.Println(" Desea volver a intentar? 1-NO  0-SI: ")
						fmt.Scan(&option2)
					} else {
						///Si lo encontro
						banderin = true ///Y aqui el ciclo que lo contiene
						break           ///cortamos el mini ciclo
					}
				}
				if option2 == 1 {
					break
				}

			}

			///Si lo encontro
			if banderin == true {
				///Si no hubo error se muestra al alumno
				mostrar_Un_Alumno(auxiliar)
				///Mostramos las opciones para el mismo

				var input string

				for option2 != 3 {
					menuAlumno()
					fmt.Scan(&input)
					num, err := strconv.Atoi(input)

					if err != nil {
						fmt.Println(" **Error: Debe ingresar un número**")
						continue
					}

					///Parseo de la opcion ingresada
					option2 = num
					switch option2 {
					case 1: ///Modificar alumno
						option3 := 0
						for option3 != 4 {
							subMenuAlumno()
							fmt.Scan(&input)
							num, err := strconv.Atoi(input)

							if err != nil {
								fmt.Println(" **Error: Debe ingresar un número**")
								continue
							}
							option3 = num
							switch option3 {
							case 1: ///Modificar nombre
								var nombre string
								fmt.Printf("\n Nuevo nombre: ")
								fmt.Scan(&nombre)
								modificarNombre(legajo, nombre, arregloAlumnos)
								fmt.Println(" Alumno actualizado correctamente ")

								break
							case 2: ///Modificar apellido
								var apellido string
								fmt.Printf("\n Nuevo nombre: ")
								fmt.Scan(&apellido)
								modificarApellido(legajo, apellido, arregloAlumnos)
								fmt.Println(" Alumno actualizado correctamente ")
								break
							case 3: ///Menu de notas
								optionMenu := 0
								for optionMenu != 4 {
									menuNotas()
									fmt.Scan(&input)
									num, err := strconv.Atoi(input)
									if err != nil {
										fmt.Println(" **Error: Debe ingresar un número**")
										continue
									}
									optionMenu = num
									switch optionMenu {
									case 1:
										var materia string
										var nuevaNota float32
										for {
											fmt.Println(" Ingrese la materia: ")
											fmt.Scan(&materia)
											fmt.Println(" Ingrese la nueva nota: ")
											fmt.Scan(&nuevaNota)
											err := modificarCalificacion(nuevaNota, arregloAlumnos, legajo, materia)
											if err != nil {
												fmt.Println(err)
												continue
											}
											break
										}
										fmt.Println(" Nota modificada exitosamente")

										break
									case 2:

										var materia string

										for {
											fmt.Println(" Ingrese la materia: ")
											fmt.Scan(&materia)
											err := eliminarCalificacion(legajo, materia, arregloAlumnos)
											if err != nil {
												fmt.Println(err)
												continue
											}
											break
										}
										fmt.Println(" Alumno dado de baja en: ", materia)
										break

									case 3:
										var materia string
										var nota float32
										for {
											fmt.Println(" Ingrese nueva materia: ")
											fmt.Scan(&materia)
											err := encontrarMateria(materia, arregloAlumnos, legajo)
											if err != nil {
												fmt.Println(err)
												continue
											}
											break
										}
										fmt.Printf("\n Ingrese nueva nota: ")
										fmt.Scan(&nota)
										agregarMateriaYnota(materia, legajo, arregloAlumnos, nota)
										fmt.Println(" Materia agregada con exito")

										break
									case 4:
										break
									default:
										fmt.Println(" Lo sentimos opcion incorrecta")
									}
								}

								break
							case 4: ///salida
								break
							default:
								fmt.Println(" Opcion incorrecta")
								fmt.Printf("\n")
							}

						}
						break

					case 2: ///Eliminar alumno
						arregloAlumnos = EliminarAlumno(arregloAlumnos, legajo)
						fmt.Println(" Alumno eliminado")

						///En caso de que lo elimine automaticamente la option se vuelve 3
						///Porque sino el submenu dice modificar alumno
						///Y se podria romper ya que no existe mas
						option2 = 3
						break
					case 3: ///Salir
						break
					default:
						fmt.Println(" Opcion Incorrecta")
					}
				}

			}

			break

		case 4: ///Si no pongo un case vacio para la opcion salir
			///Se toma como un default y dice opcion incorrecta
			break
		default:
			fmt.Println(" Opcion incorrecta")
		}
	}

	fmt.Println(" Programa finalizado")
}

func menuPrincipal() {
	fmt.Println(" _ _ _ _ _ _ _ _ _  ")
	fmt.Println("|   Menu Princial  |")
	fmt.Println("| 1-Cargar Alumnos |")
	fmt.Println("| 2-Ver Alumnos    |")
	fmt.Println("| 3-Buscar Alumno  |")
	fmt.Println("| 4-Salir          |")
	fmt.Println(" ------------------ ")
	fmt.Printf(" Para modificar notas pulse la opcion 3\n")
	fmt.Printf(" Pulse la opcion que desee:  ")
}
func CargaAlumnos(alumnos []Alumno) []Alumno {

	///Variables para utilizar en la funcion
	var auxiliar Alumno
	var legajoA string
	var seguir int
	///***********************************///

	for seguir != 1 {

		fmt.Print("\n Legajo (Alfanumerico): ")
		fmt.Scan(&legajoA)
		err := legajoExistente(legajoA, alumnos)
		if err != nil {
			fmt.Println(err)
			continue
		} else {
			auxiliar.legajo = legajoA
		}

		fmt.Printf("\n Nombre: ")
		fmt.Scan(&auxiliar.nombre)
		if auxiliar.nombre == "" {
			fmt.Printf("Ingrese un nombre")
			continue
		}

		fmt.Printf("\n Apellido: ")
		fmt.Scan(&auxiliar.apellido)

		fmt.Printf("\n ****Carga de materias y de notas****")
		auxiliar.materiaCalificacion = cargaDeMateriasYnotas()

		idAutoincremental++             ///Incrementa la variable global
		auxiliar.id = idAutoincremental ///La guardamos en el alumno

		alumnos = append(alumnos, auxiliar)
		fmt.Printf("\n Continuar cargando alumnos?(0-SI  1-NO) --> ")
		fmt.Scan(&seguir)
	}

	///Modificamos la variable global
	///Ahora si se podran ver los alumnos
	alumnosCargados = true

	return alumnos

}
func cargaDeMateriasYnotas() map[string]float32 {

	//Mapa auxiliar y variables
	calificaciones := make(map[string]float32)
	seguir := 0
	var materia string
	var notaAux string ///Esta va ser parseada
	var nota float32
	//******************************

	for seguir != 1 {

		fmt.Printf("\n Materia: ")
		fmt.Scan(&materia)
		fmt.Printf("\n Calificacion: ")
		fmt.Scan(&notaAux)
		num, err := strconv.ParseFloat(notaAux, 32)

		if err != nil {
			fmt.Println(" **Error: Debe ingresar un número**")
			continue
		}
		nota = float32(num)

		calificaciones[materia] = nota

		fmt.Printf("\n Continuar cargando materias?(0-SI  1-NO) --> ")
		fmt.Scan(&seguir)
	}
	return calificaciones
}
func legajoExistente(legajo string, alumnos []Alumno) error {

	for i := 0; i < len(alumnos); i++ {
		if alumnos[i].legajo == legajo {
			return errors.New(" Legajo existente, cargue otro")
		}
	}
	return nil
}
func muestraAlumnos(alumnos []Alumno) {
	for i := 0; i < len(alumnos); i++ {
		mostrar_Un_Alumno(alumnos[i])
	}
}

// /Funcion modularizada para mostrar un alumno
func mostrar_Un_Alumno(aux Alumno) {

	fmt.Printf("\n|Alumno Id:%d        ", aux.id)
	fmt.Printf("\n|Legajo:%s           ", aux.legajo)
	fmt.Printf("\n|Nombre:%s           ", aux.nombre)
	fmt.Printf("\n|Apellido:%s         \n", aux.apellido)
	mostrar_Materias_Notas(aux.materiaCalificacion)
}

func mostrarAlumnoSinNotas(aux Alumno) {
	fmt.Printf("\n|Alumno Id:%d        ", aux.id)
	fmt.Printf("\n|Legajo:%s           ", aux.legajo)
	fmt.Printf("\n|Nombre:%s           ", aux.nombre)
	fmt.Printf("\n|Apellido:%s         ", aux.apellido)
}

func mostrar_Materias_Notas(mapita map[string]float32) {
	fmt.Println(" _ _ _ _ _ _ _ _ _")
	fmt.Println(" Estado Academico ")
	for materia, califiacion := range mapita {
		fmt.Printf("\n %s: %.2f\n", materia, califiacion)
	}
	fmt.Println(" ------------------")
}
func buscarAlumno(legajo string, arregloAlumnos []Alumno) (error, Alumno) {

	var auxiliar Alumno
	encontrado := false

	for i := 0; i < len(arregloAlumnos); i++ {
		if legajo == arregloAlumnos[i].legajo {
			auxiliar = arregloAlumnos[i]
			encontrado = true
			break
		}
	}

	if encontrado == false {
		return errors.New(" El legajo no existe"), auxiliar
	}

	return nil, auxiliar

}
func menuAlumno() {
	fmt.Println("  _ _ _ _ _ _ _ _ _ _  ")
	fmt.Println(" |   Menu del alumno  |")
	fmt.Println(" |1- Modificar alumno |")
	fmt.Println(" |2- Eliminar alumno  |")
	fmt.Println(" |3- Salir            |")
	fmt.Println("  -------------------- ")
}
func subMenuAlumno() {
	fmt.Println("  _ _ _ _ _ _ _ _ _ _ _  ")
	fmt.Println(" |  Opciones del alumno |")
	fmt.Println(" |1- Modificar nombre   |")
	fmt.Println(" |2- Modificar apellido |")
	fmt.Println(" |3- Gestionar notas    |")
	fmt.Println(" |4- Salir              |")
	fmt.Println("  ---------------------- ")

}
func menuCalificacionesAlumno() {
	fmt.Println(" |1- Modificar una nota |")
	fmt.Println(" |2- Agregar una nota   |")
	fmt.Println(" |3- Eliminar una nota  |")
	fmt.Println(" |4- Salir              |")
}
func EliminarAlumno(arregloAlumnos []Alumno, legajo string) []Alumno {

	for i, alumnito := range arregloAlumnos {
		if alumnito.legajo == legajo {
			arregloAlumnos = append(arregloAlumnos[:i], arregloAlumnos[i+1:]...)
		}
	}
	return arregloAlumnos
}
func modificarNombre(legajo, nombre string, arregloAlumnos []Alumno) {

	for i := 0; i < len(arregloAlumnos); i++ {
		if arregloAlumnos[i].legajo == legajo {
			arregloAlumnos[i].nombre = nombre
			break
		}
	}
}
func modificarApellido(legajo, apellido string, arregloAlumnos []Alumno) {

	for i := 0; i < len(arregloAlumnos); i++ {
		if arregloAlumnos[i].legajo == legajo {
			arregloAlumnos[i].apellido = apellido
			break
		}
	}
}
func menuNotas() {
	fmt.Println("  _ _ _ _ _ _ _ _ _ _ _ _ _  ")
	fmt.Println(" |   Menu gestion de notas  |")
	fmt.Println(" |1- Modificar calificacion |")
	fmt.Println(" |2- Dar de baja materia    |")
	fmt.Println(" |3- Agregar materia y nota |")
	fmt.Println(" |4- Salir                  |")
	fmt.Println("  --------------------------")

}

func modificarCalificacion(nuevaNota float32, arreglo []Alumno, legajo, materia string) error {
	for i := 0; i < len(arreglo); i++ {
		if arreglo[i].legajo == legajo {
			for nombreMateria := range arreglo[i].materiaCalificacion {
				if nombreMateria == materia {
					arreglo[i].materiaCalificacion[nombreMateria] = nuevaNota
					return nil
				}
			}
		}
	}
	return errors.New(" La materia ingresada no existe")

}

func eliminarCalificacion(legajo, materia string, arreglo []Alumno) error {

	for i := 0; i < len(arreglo); i++ {
		if arreglo[i].legajo == legajo {
			for nombreMateria := range arreglo[i].materiaCalificacion {
				if nombreMateria == materia {
					delete(arreglo[i].materiaCalificacion, materia)
					return nil
				}
			}
		}
	}

	return errors.New(" La materia no existe")
}

func encontrarMateria(materia string, arreglo []Alumno, legajo string) error {
	for i := 0; i < len(arreglo); i++ {
		if arreglo[i].legajo == legajo {
			for nombreMateria := range arreglo[i].materiaCalificacion {
				if nombreMateria == materia {
					return errors.New(" La materia ya existe")
				}
			}
		}
	}
	return nil
}

func agregarMateriaYnota(materia, legajo string, arreglo []Alumno, nota float32) {

	for i := 0; i < len(arreglo); i++ {
		if arreglo[i].legajo == legajo {
			arreglo[i].materiaCalificacion[materia] = nota
			break
		}
	}

}
func promedios(arreglo []Alumno) {

	for i := 0; i < len(arreglo); i++ {
		mostrarAlumnoSinNotas(arreglo[i])
		promedio := calculcarPromedio(arreglo[i].materiaCalificacion)
		fmt.Printf("\n|Promedio general: %.2f\n", promedio)
	}
}

func calculcarPromedio(alumnoNotas map[string]float32) float32 {

	var promedio float32
	var acumulador float32
	for _, nota := range alumnoNotas {
		acumulador += nota
	}

	promedio = acumulador / (float32(len(alumnoNotas)))
	return promedio
}
