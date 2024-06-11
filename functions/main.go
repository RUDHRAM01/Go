package main

import "fmt"

func main() {
	greetings()

	sum := adder(3, 5)
	fmt.Println(sum)

	sum2, _ := advanceAdder(1, 2, 3, 4, 5)
	fmt.Println(sum2)
}

func greetings() {
	fmt.Println("Hello, World!")
}

func adder(a int, b int) int {
	return a + b
}

func advanceAdder(values ...int) (int, string) {
	sum := 0
	for _, value := range values {
		sum += value
	}
	return sum, "done"
}
