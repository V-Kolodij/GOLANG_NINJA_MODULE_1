package main

import "fmt"

var msg string

func init() {
	msg = "init func"
}

func main() {
	fmt.Println(msg)
}
