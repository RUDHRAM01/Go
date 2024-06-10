package main

import "fmt"

func main() {
	var arr = []int{1, 2, 3, 4, 5}
	fmt.Println(arr)

	arr = append(arr, 6)
	fmt.Println(arr)

	fmt.Println(arr[1:3])

	// dynamic array

	var arr2 = make([]int, 5)
	fmt.Println(arr2)

	fmt.Println((append(arr[:1], arr[2:]...)))
}
