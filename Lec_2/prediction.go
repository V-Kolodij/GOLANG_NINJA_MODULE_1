package main

import "fmt"
import "errors"

func main() {
	fmt.Println(prediction("monday"))
}

func prediction(dayOfWeek string) (string, error) {
	switch dayOfWeek {
	case "monday":
		return "Melone", nil
	case "tuesday":
		return "Kukaracha", nil
	default:
		return "Unknown day of week", errors.New("invalid day of week")
	}
}
