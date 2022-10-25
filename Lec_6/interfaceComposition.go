package main

import (
	"fmt"
	"math"
)

type Shape2 interface {
	ShapeWithArea
	ShapeWithPerimeter
}

type ShapeWithArea interface {
	Area() float32
}
type ShapeWithPerimeter interface {
	Perimeter() float32
}
type Square2 struct {
	sideLength float32
}

func (s Square2) Area() float32 {
	return s.sideLength * s.sideLength
}

func (s Square2) Perimeter() float32 {
	return s.sideLength * 2
}

type Circle2 struct {
	radius float32
}

func (c Circle2) Area() float32 {
	return c.radius * c.radius * math.Pi
}

func main() {
	square := Square2{4}
	//circle := Circle2{5}

	printShapeArea2(square) // work, Square2 implements methods from two interfaces
	//printShapeArea2(circle) // not work because not implements method Perimeter

}

func printShapeArea2(shape Shape2) {
	fmt.Println(shape.Area())
	fmt.Println(shape.Perimeter())
}
