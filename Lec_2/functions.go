package main

import "fmt"

func main() {
	//printMessage("Call 1")
	//printMessage("Call 2")
	//printMessage("Call 3")

	message := sayHello("Vasya", 25)
	printMessage(message)
}

func printMessage(message string) {
	fmt.Println(message)
}

func sayHello(name string, age int) string {
	return fmt.Sprintf("Hello, %s, You are %d years old))", name, age)
}
