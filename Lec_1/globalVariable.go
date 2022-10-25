package main

import "fmt"

var a, b, c int

func main() {
	a, b, c := 1, 2, 3
	fmt.Println(a, b, c)
	printGlobalVariables()
}

func printGlobalVariables() {
	fmt.Println(a, b, c)
}
