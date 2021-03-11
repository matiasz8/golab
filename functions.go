package main

import (
	"fmt"
)

// User is the structure for commercial area
type User struct {
	age           int
	name, surname string
}

func (user User) getUser() string {
	return user.name + " " + user.surname
}

func (user User) setUserOnContext(n string) {
	user.name = n
	fmt.Println("User changed on context: " + user.name)
}

func (user *User) setUserThoughVal(n string) {
	user.name = n
	fmt.Println("User name changed accesing by memory direction: " + user.name)
}

func main() {
	person1 := new(User)
	person1.name = "John"
	person1.surname = "Smith"
	fmt.Println(person1.name)
	fmt.Println(person1.getUser())

	person1.setUserOnContext("Felix")
	fmt.Println("Actual value of User.name: " + person1.name)

	person1.setUserThoughVal("Peter")
	fmt.Println("Actual value of User.name: " + person1.name)

}
