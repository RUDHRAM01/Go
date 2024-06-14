package main

import "fmt"

func main() {
	rudharm := Student{
		"rudhram@gmail.com",
		"Rudhram",
		22,
	}

	fmt.Println(rudharm.GetEmail())
}

type Student struct {
	Email string
	Name  string
	Age   int
}

func (s Student) GetEmail() string {
	return s.Email
}
