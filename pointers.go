package main

import (
	"fmt"
)

func main() {
	/*
		punteros:
		1- Es una dir de mem
		2- En lugar de guardar el valor, tenemos la dir donde está el valor
		3- x,y => 0Xadedfeged => 6
		4- X => 0Xadedfeged => 6
		5- Y => 0Xadedfeged => 6
		*T => *int, *string, *Struct
		Valor zero => nil
	*/

	var x, y *int

	mark := 1

	x = &mark
	y = &mark

	// altero el valor
	*x = 9

	fmt.Println(x, *x)
	fmt.Println(y, *y)
}
