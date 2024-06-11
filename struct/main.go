package main

import "fmt"

func main() {
	rudharm := Student{
		"rudhram@gmail.com",
		"Rudhram",
		22,
	}

	fmt.Printf("rudhram's email %v and age %v\n", rudharm.Email, rudharm.Age)
	fmt.Printf("full details of rudhram %+v", rudharm)
}

type Student struct {
	Email string
	Name  string
	Age   int
}
