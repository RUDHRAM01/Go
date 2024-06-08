package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	value, _ := reader.ReadString('\n')

	number, err := strconv.ParseFloat(strings.TrimSpace(value), 64)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	} else {
		fmt.Println("Number: ", number+1)
	}
}
