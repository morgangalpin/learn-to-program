package main

import (
	"fmt"
	"os"
)

func main() {
	// Args is the name of the array of command line arguments.
	billy := os.Args[1]
	likes := os.Args[2]

	if billy == "Goat" {
		if likes == "flowers" {
			fmt.Println("Billy is a goat and likes flowers.")
		} else {
			fmt.Println("Billy is a goat but doesn't like flowers.")
		}
	} else {
		if likes == "flowers" {
			fmt.Println("Billy isn't a goat but likes flowers.")
		} else {
			fmt.Println("Billy isn't a goat and doesn't like flowers.")
		}
	}
}
