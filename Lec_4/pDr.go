package main

import "fmt"

// panic deffer recover
//func main() {
//	panic("aaaaaaaaaaa, help!")
//}

//Defer call printMassage in end
//func main() {
//	defer printMessage()
//	fmt.Println("main()")
//	fmt.Println("main() 2")
//	fmt.Println("main() 3")
//}
//
//func printMessage() {
//	fmt.Println("func printMessage()")
//}
//

//HANDLE PANIC

func main() {
	defer handlerPanic()
	messages := []string{
		"message 1",
		"message 2",
		"message 3",
		"message 4",
	}
	messages[4] = "message 5"

	fmt.Println(messages)
}

func handlerPanic() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}
	fmt.Println("handler panic success")
}
