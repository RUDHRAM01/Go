package main

import (
	"fmt"
	"time"
)

func main() {
	time := time.Now()

	fmt.Println("Hello, World! The time is: ", time.Format("01-02-2006 Monday"))
}
