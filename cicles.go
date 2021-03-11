package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println("Iteration: ", i)
	}

	j := 0
	for j < 5 {
		println("Just declared limit: ", j)
		j++
	}

	k := 0
	for {
		println("For without params: ", k)
		k++
		if k == 5 {
			break
		}
	}

	l := 0
	for {
		if l == 2 {
			l++
			continue
		}
		println(l)
		l++

		if l == 5 {
			break
		}

	}

	sharks := []string{"hammerhead", "great white", "dogfish"}

	for i, shark := range sharks {
		fmt.Println(i, shark)
	}

}

//https://www.digitalocean.com/community/tutorials/how-to-construct-for-loops-in-go-es
