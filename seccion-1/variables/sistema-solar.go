package main
import "fmt"

func main() {
	lunas := [8]int{0, 0, 1, 2, 79, 82, 27, 14}
	planetas := [8]string{"Mercurio", "Venus", "Tierra", "Marte", "Jupiter", "Saturno", "Urano", "Neptuno"}
	j := 0
	for i := 0; i < len(planetas); i++ {
		fmt.Printf("%s tiene %d lunas \n", planetas[j], lunas[i])
		j++
	}
}