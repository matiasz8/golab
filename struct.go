package main

import (
	"fmt"
)

type User2 struct {
	age           int
	name, surname string
}

func main() {
	var person1 User
	fmt.Println(person1)

	fmt.Println(User{name: "Tom", surname: "Phil"})

	person2 := User{name: "Jerry", surname: "Louis"}
	fmt.Println(person2)

	person3 := User{22, "Tim", "Jonas"}
	fmt.Println(person3)

	person4 := new(User)
	person4.name = "Gary"
	fmt.Println(person4)
	fmt.Println((*person4).name)
	fmt.Println(person4.name)
}
