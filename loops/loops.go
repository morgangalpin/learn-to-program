package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var str string = "1"
	var number int = 1
	numberOfGreetings, _ := strconv.Atoi(os.Args[1])
	for i := 0; i < numberOfGreetings; i++ {
		fmt.Println("Hello!")
		fmt.Println("Goodbye!")
	}
}
