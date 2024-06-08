package main

import "fmt"

const LoginToken string = "asdasdasd" // HERE capital letter means public variable

func main() {
	var username string = "root"
	var isMan bool = true

	fmt.Println(username)
	fmt.Println(isMan)

	// we can also declare without specifying the type
	var name = "root"
	var isMan2 = true
	fmt.Println(name)
	fmt.Println(isMan2)

	// we can also declare without using var keyword
	email := "test@gmail.com" // this way is only available inside a function
	isMan3 := false
	fmt.Println(email)
	fmt.Println(isMan3)

	fmt.Println(LoginToken)
	fmt.Printf("Login Token Type: %T", LoginToken)
}
