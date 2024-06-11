package main

func main() {
	age := 12

	if age > 18 {
		println("You are an adult")
	} else {
		println("You are a minor")
	}

	// we can create a variable in the if statement
	if age := 22; age > 18 {
		println("You are an adult")
	} else {
		println("You are a minor")
	}
}
