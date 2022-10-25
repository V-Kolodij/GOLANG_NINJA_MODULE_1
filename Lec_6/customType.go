package main

import (
	"fmt"
)

//Create custom type (for custom type we can add methods)

type Age int

func (a Age) isAdult() bool {
	return a >= 18
}

type User struct {
	name   string
	age    Age
	sex    string
	weight int
	height int
}

func newUser(name, sex string, age, weight, height int) User {
	return User{
		name:   name,
		age:    Age(age),
		sex:    sex,
		weight: weight,
		height: height,
	}
}

func main() {
	user1 := newUser("Kolya", "Male", 23, 87, 178)
	fmt.Println(user1.age.isAdult())
}
