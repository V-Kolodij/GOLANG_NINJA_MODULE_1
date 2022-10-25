package main

import "fmt"

func main() {
	users := map[string]int{
		"Vasya": 15,
		"Kolya": 23,
		"Petya": 18,
	}

	age, exists := users["Kolya"]

	if exists {
		fmt.Println(age)
		fmt.Println(exists)
	}

	fmt.Println("no in list")
	//fmt.Println(users)
	//fmt.Println(users["Vasya"])

	for key, value := range users {
		fmt.Println(key)
		fmt.Println(value)
	}
	//Add value
	users["Sergey"] = 21
	fmt.Println(users)

	//delete

	delete(users, "Vasya")

	//Make map for add
	var users2 map[string]int      // not add element (panic -> map nil)
	users3 := make(map[string]int) // can add element

	fmt.Println(users2)
	users["Dimon"] = 45
	fmt.Println(users3)

}
