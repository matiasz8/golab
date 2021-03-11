package main

import (
	"fmt"
)

func main() {
	// var arrayInt [10]int
	// var arrayStr [10]string

	arrayInt := [4]int{1, 2, 4, 56}
	arrayStr := [10]string{}

	fmt.Println(arrayInt)
	fmt.Println(arrayStr)

	for i := 0; i < len(arrayInt); i++ {
		println(arrayInt[i])

	}

	var matrix [3][2]int
	matrix[0][1] = 1

	fmt.Println(matrix)
}
