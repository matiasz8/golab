package main

import (
	"fmt"
)

func main() {
	// inicializo el slice en 0
	slice := make([]int, 3, 5)
	fmt.Println(slice)

	slice = append(slice, 2)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}
