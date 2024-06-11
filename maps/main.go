package main

import "fmt"

func main() {
	languages := make(map[string]string)

	languages["en"] = "English"
	languages["id"] = "Indonesia"
	languages["jp"] = "Japan"

	fmt.Println(languages)

	delete(languages, "jp")
	fmt.Println(languages)

	for key, value := range languages {
		fmt.Println(key, value)
	}
}
