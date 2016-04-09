package main

import "fmt" //Pemite cargar el print, el scan

/*func main() { //Genera la funcion 
  fmt.Print("Ingrese un numero: ")
  var input float64 //Variable llamada input tipo float
  fmt.Scanf("%f", &input) //guarda el valor que entro el usuario en la variable

  output := input * 2 //le multiplica 2 a la variable

  fmt.Println(output)
}*/


/*EJERCICIO: Cambie el anterior programa para que en vez de capturar
	un decimal, capture un texto y lo imprima.*/

	//Ejercicio
 func main() { //Genera la funcion 
  fmt.Print("Ingrese un texto: ")
  var input string //Variable llamada input tipo float
  fmt.Scanf("%s", &input) //guarda el valor que entro el usuario en la variable


  fmt.Println(input)
}