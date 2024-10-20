package main

import "fmt"

var a int = 20
var b string = "nome"

func main() {

	var c float64 = float64(a) // conversão de tipo int em float
	fmt.Printf("O Valor de c é %g %T o valor de b é %s %T \n", c, c, b, b)

}
