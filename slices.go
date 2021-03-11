package main

import (
	"fmt"
)

func main() {
	valueSlice := []int{1, 33, 44, 66, 22, 4, 7}
	fmt.Println(valueSlice)

	valueSlice2 := []int{1, 28}
	if valueSlice2 == nil {
		fmt.Println("value empty")
	} else {
		fmt.Println(valueSlice2)
	}

	// slicing
	array := [3]int{1, 7, 5}
	slice := array[:2]
	fmt.Println(slice)

}
