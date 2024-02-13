package main

import (
	"fmt"
	"slices"
)

type Pila struct {
	cantidad int
	valores  []int
}

func (p *Pila) Add(valor int) {
	p.valores = append(p.valores, valor)
	p.cantidad += 1
}

func (p *Pila) RemoveAt(index int) bool {
	banderita := true
	booleanito := p.isInBoundaries(index)
	if !booleanito {
		banderita = false
	} else {
		p.valores = slices.Delete(p.valores, index-1, index)
	}
	p.cantidad -= 1
	return banderita
}

// /Si manejamos un slice o un arreglo la cantidad de elementos de la struct la podemos utilizar con los metodos de los mismos (como el len)
func (p Pila) Lenght() int {
	cantidad := len(p.valores)
	return cantidad
}

func (p *Pila) Set(index, valor int) bool {
	banderita := true
	booleanito := p.isInBoundaries(index)

	if !booleanito {
		banderita = false
	} else {
		p.valores[index-1] = valor
	}
	return banderita
}

func (p Pila) ToString() {
	fmt.Printf("\n--------------------\n")
	for i := 0; i < len(p.valores); i++ {
		fmt.Printf(" [%d] ", p.valores[i])
	}
	fmt.Printf("\n--------------------\n")
}

/*Funcion extra devuelve si el indice ingresado esta dentro de la longitud del slice
Supongamos que tenemos 6 elementos en el slice, el indice 6 no existe, pero el usuario no lo sabe
Por lo que si ingresa el 6 (refiriendose a la posicion [5]) diremos que esta dentro de la longitud
y en la funcion del set solo debemos restar uno a la posicion ingresada
*/

func (p Pila) isInBoundaries(index int) bool {
	return index <= p.cantidad
}

func main() {

	var pilin Pila
	pilin.Add(1)
	pilin.Add(2)
	pilin.Add(3)
	pilin.Add(4)
	pilin.Add(5)
	pilin.Add(6)

	pilin.ToString()

	pilin.RemoveAt(6)
	pilin.Set(1, 8)

	pilin.ToString()

}
