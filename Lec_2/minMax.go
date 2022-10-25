package main

import "fmt"

func main() {
	fmt.Println(findMin(3, 4, 5, 6, 8, 9, 0, -5))
	fmt.Println(findMax(23, 34, 23, 456, 676, 45, 3, 23, 45, 0))
}
func findMin(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}
	min := numbers[0]
	for _, i := range numbers {
		if i < min {
			min = i
		}
	}

	return min
}

func findMax(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}
	max := numbers[0]
	for _, i := range numbers {
		if i > max {
			max = i
		}
	}

	return max
}
