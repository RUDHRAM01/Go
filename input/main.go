package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// take input from user
	reader := bufio.NewReader(os.Stdin)
	value, _ := reader.ReadString('\n')

	fmt.Println("You entered: ", value)

}
