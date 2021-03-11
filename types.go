package main

import (
	"fmt"
	"strconv"
)

func main() {
	age := 22
	cuit := "203"
	fmt.Println(age)

	ageStr := strconv.Itoa(age)
	cuitInt, _ := strconv.Atoi(cuit)
	fmt.Println("My age is " + ageStr)
	fmt.Println(cuitInt + 5)
}
