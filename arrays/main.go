package main

import "fmt"

func main() {
	var numbersArray [5]int
	numbersArray[0] = 1

	// length of the array
	fmt.Println(len(numbersArray))

	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)
}
