package main

import "fmt"

/*func main() {
	var x [5]float64 //Creo el arreglo de 5 posiciones y tipo float
	x[0] = 98 //Inicializo
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 38

	var total float64 = 0 //Declaracion e inicializacion var Total
	for i := 0; i < 5; i++ {
		total += x[i] //suma todos los elementos del arreglo
	}
	fmt.Println(total / 5) //Muestra el promedio
}*/

/*EJERCICIO: Escribir un programa que encuentre el numero menor en el 
siguiente arreglo y lo imprima: */


func main() {
	x := []int{
	  48,96,86,68,
	  57,82,63,70,
	  37,34,83,27,
	  19,97, 9,17,
	}

	var menor int = 0 //Declaracion e inicializacion var menor
	menor = x[0]
	for i := 0; i < 15; i++ {		
		if x[i] < menor { //Condicional
			menor = x[i]
		}		
	}
	fmt.Printf("El numero menor es: %d\n", menor) //Muestra el menor
}