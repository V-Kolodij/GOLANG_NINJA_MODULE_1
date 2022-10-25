package main

import "fmt"

func main() {
	a := []int64{1, 2, 3, 4}
	b := []float64{1.1, 2.2, 3.3}
	c := []string{"1", "2", "3"}

	fmt.Println(sum(a))
	fmt.Println(sum(b))
	fmt.Println(sum2(b))
	fmt.Println(searchEl(c, "2"))

}

func sum[V int64 | float64](input []V) V {
	var result V
	for _, number := range input {
		result += number
	}
	return result
}

// OR with interface:
type Number interface {
	int64 | float64
}

func sum2[V Number](input []V) V {
	var result V
	for _, number := range input {
		result += number
	}
	return result
}

// comparable : all types or struct without slices and maps
func searchEl[C comparable](elements []C, search C) bool {
	for _, el := range elements {
		if el == search {
			return true
		}
	}

	return false
}

//Any interface (any types)

func printAny[A any](input A) {
	fmt.Println(input)
}
