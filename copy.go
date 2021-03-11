package main

import (
	"fmt"
)

func main() {
	/*
		copy, copia el mínimo de elementos
	*/
	slice := []int{1, 3, 55, 2, 33, 9, 7}
	// impostor := make([]int, 6)
	impostor := make([]int, len(slice), cap(slice)*2)
	copy(impostor, slice)
	fmt.Println(slice)
	fmt.Println(impostor)
}
