package main

import (
	"fmt"
	"reflect"
)

func main() {
	message := "I will become a Golang Ninja :))"
	fmt.Println(message)
	fmt.Println(reflect.TypeOf(message))
	a, b, c := 1, 2, 3
	a, b = b, a
	fmt.Println(a, b, c)
	var bYte byte = 95
	fmt.Printf("%c\n", bYte)
	var rUne rune = '~'
	var rune2 rune = 164
	fmt.Println(rUne)
	fmt.Printf("Message: %s, Rune: %c, Rune2: %c", message, rUne, rune2)
}
