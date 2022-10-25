package main

import (
	"errors"
	"fmt"
)

//func main() {
//	message := "I am ninja"
//	printMessage(message)
//	fmt.Println(message)
//}
//
//func printMessage(message string) {
//	message += ", Golang Ninja ))) "
//	fmt.Println(message)
//}
//
//func main() {
//	message := "I am ninja"
//	fmt.Println(message)
//	changeMessage(&message)
//	fmt.Println(message)
//}
//
//func changeMessage(message *string) {
//	*message += ", Golang Ninja ))) "
//}

//func main() {
//	messages := [3]string{"1", "2", "3"} // array
//	messages[1] = "5"
//	fmt.Println(messages)
//	fmt.Println(messages[2])
//	printMessage(messages)
//}
//
//func printMessage(messages [3]string) error {
//	if len(messages) == 0 {
//		return errors.New("array is empty")
//	}
//	messages[1] = "9"
//	fmt.Println(messages)
//	return nil
//}

func main() {
	messages := []string{"1", "2", "3"} // slice
	printMessage(messages)
	fmt.Println(messages)
	secondSlice()
}

func printMessage(messages []string) error {
	if len(messages) == 0 {
		return errors.New("slice is empty")
	}
	messages[1] = "9"
	return nil
}

func makeSlice() {
	messages := make([]string, 5) // slice len 5 capacity 5
	fmt.Println(len(messages))
	fmt.Println(cap(messages))
	// append 6 and cap will be 10
	messages = append(messages, "6")
	fmt.Println(len(messages))
	fmt.Println(cap(messages))
}

func secondSlice() {
	matrix := make([][]int, 10)
	for x := 0; x < 10; x++ {
		matrix[x] = make([]int, 10)
		for y := 0; y < 10; y++ {
			matrix[x][y] = x
		}
		fmt.Println(matrix[x])
	}
	fmt.Println(matrix)
}

func rangeSlice() {
	messages := []string{
		"message 1",
		"message 2",
		"message 3",
		"message 4",
	}

	for i := range messages {
		fmt.Println(messages[i])
	}

	for index, value := range messages {
		fmt.Println(index)
		fmt.Println(value)
	}

	for _, message := range messages {
		fmt.Println(message)
	}

	//infinit loop
	counter := 0
	for {
		if counter == 100 {
			break
		}
		counter++
		fmt.Println(counter)
	}
}
