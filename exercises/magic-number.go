package main

import (
	"fmt"
	"os"
	"strconv"
)

// Write code to get the user to guess the magic number.
// Ask the user to pic a number between 1 and 10.
// If the user gets the number wrong, give them a hint (higher, lower)
// If the get it correct, say something nice.

func main() {
	userGuess, _ := strconv.Atoi(os.Args[1])
	fmt.Println("Your guess is", userGuess)

	number := 8
}
