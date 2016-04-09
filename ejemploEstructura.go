package main

import "fmt"


/*type Circle struct{
	x, y, r float64
}

func main() {

	var c Circle //Definir la variable
	c:= new(Circle) 
	c := Circle(0,0,5) //inicializar
}*/


type persona struct{
	nombre string
	estaura float64
}

func main() {
	milena := persona{"Mile Cardenas", 1.58}

	fmt.Printf("%s Corriendo\n", milena.nombre)
	
}