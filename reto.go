package main

import "fmt"

func main() {
  fmt.Print("Ingrese una palabra: ")

  var txt string
  fmt.Scanf("%s", &txt)

  for i:= 0; i<len(txt) ; i++{
   	fmt.Println(txt[:i+1])
  }

  for i:= len(txt)-1; i>0 ; i--{
  	fmt.Println(txt[:i+1])
  }
}


