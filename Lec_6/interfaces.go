package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float32
}

type Square struct {
	sideLength float32
}

func (s Square) Area() float32 {
	return s.sideLength * s.sideLength
}

type Circle struct {
	radius float32
}

func (c Circle) Area() float32 {
	return c.radius * c.radius * math.Pi
}

func main() {
	square := Square{4}
	circle := Circle{5}

	printShapeArea(square)
	printShapeArea(circle)

	printInterface(circle)
	printInterface(square)
	printInterface("qwert")
	printInterface(789)

	printTypeInterface(678)
	printTypeInterface(circle)

	d, _ := fmt.Println("12345")
	fmt.Println(d)

	typeLead("SDVASDV")
	typeLead(22)
}

func printShapeArea(shape Shape) {
	fmt.Println(shape.Area())
}

//Empty interface

func printInterface(i interface{}) {
	fmt.Printf("%+v\n", i)
}

func printTypeInterface(i interface{}) {
	switch value := i.(type) {
	case int:
		fmt.Println("int", value)
	case bool:
		fmt.Println("bool", value)
	default:
		fmt.Println("Unknown Type")
	}
}

func typeLead(i interface{}) {
	str, ok := i.(string)
	if !ok {
		fmt.Println("interface is not string")
		return
	}

	fmt.Println(len(str))
}
