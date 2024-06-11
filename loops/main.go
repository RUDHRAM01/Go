package main

import "fmt"

func main() {
	name := []string{"John", "Doe", "Smith"}
	// for i := 0; i < len(name); i++ {
	// 	fmt.Println(name[i])
	// }

	// for i := range name {
	// 	fmt.Println(name[i])
	// }

	// for key, value := range name {
	// 	fmt.Println(key, value)
	// }

	var index int = 0
	for index < len(name) {
		fmt.Println(name[index])
		index++
	}
}
