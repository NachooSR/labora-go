package main
import ("fmt"
 "os"
 "bufio"
 "strings"
)

func main(){

	reader:= bufio.NewReader(os.Stdin)
	fmt.Printf("Nombre del archivo: ")
	
    
	texto,err:= reader.ReadString('\n')
	 if(err!=nil){
	 fmt.Println("ERROR")
	 return
	}
	
	texto=strings.TrimSpace(texto)

	archivo,err:=os.Open(texto)
	if(err!=nil){
		fmt.Println("Error al abrir el archivo  no existe")
	   return
	}
	
	///C:\Users\silva\OneDrive\Escritorio\fotocv.png

	defer archivo.Close()
	fmt.Println("No hubo error")
}
