package main

import "fmt"

//Los numeros como telefonos, o dnis, con los cuales no voy a operar matematicamente los almaceno como string para no utilizar un numero tan largo

type Persona struct {
	nombre   string
	edad     int
	ciudad   string
	telefono string
}

func(p *Persona)cambiarCiudad(ciudad string){
   if(p.ciudad!=ciudad){
       p.ciudad=ciudad
   }
}

func main() {
	
	personita := Persona{
		nombre:   "Carlos",
		edad:     25,
		ciudad:   "Mar del Plata",
		telefono: "2235662067",
	}
	personita2 := Persona{
		nombre:   "Carla",
		edad:     20,
		ciudad:   "Otamendi",
		telefono: "2234229607",
	}

	mostrarPersona(personita)
	mostrarPersona(personita2)

	personita.cambiarCiudad("Carlos Paz")
	fmt.Println("Valor actualizado")
	mostrarPersona(personita)

}


func mostrarPersona(personita Persona) {
	fmt.Println("\n-----------")
	fmt.Println("Nombre: ", personita.nombre)
	fmt.Println("Edad: ", personita.edad)
	fmt.Println("Ciudad: ", personita.ciudad)
	fmt.Println("Nro Tel: ", personita.telefono)
	fmt.Println("-----------\n")
}
