package main

import (
	"fmt"
	"os"
	"strconv"
)

// Cookies for Timmy and Suzie
// 
// Little Timmy has some cookies. How many he starts with is up to
// the user. The Timmy starts giving cookies to his friend Suzie.
// At some point, Timmy will finish giving cookies to Suzie.
//
// The first command line variable is the number of cookies that 
// Timmy starts with. Then ask the user how many cookies Timmy should 
// give to Suzie. Keep asking the user how many cookies Timmy should
// give to Suzie until the user answers "0" or Timmy runs out of 
// cookies. When Timmy is done, print out how many cookies Timmy 
// gave to Suzie, and how many cookies he has left.

func main() {
	numberOfCookies, _ := strconv.Atoi(os.Args[1])
}
