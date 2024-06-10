package main

import "fmt"

func main() {
	var value = 5
	var pointer *int = &value

	fmt.Println("Value: ", pointer)
	fmt.Println("Value: ", *pointer)
}
