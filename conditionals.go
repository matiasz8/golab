package main

import (
	"fmt"
)

func main() {
	/*
		operators:
		==
		!=
		<
		>
		<=
		>=
		&&
	*/

	x := 12
	y := 15

	if x > y {
		fmt.Printf("%d in great that %d\n", x, y)
	} else if x == y {
		fmt.Println("Values evaluated are the same")
	} else {
		fmt.Printf("%d in lower that %d\n", x, y)
	}

	//https://codingornot.com/05-go-to-go-operadores
}
