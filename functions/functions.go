package main

import (
	"fmt"
)

// Functions
// 
// Functions start with the keyword `func`, then the name of the function.
// A function may or may not return a value. If it returns no value, it is
// actually returning a `nil` value.
//
// Example of a function that returns no value:
// 	func functionThatReturnsNil() {
//     fmt.Println("This is what this function does. I just prints something.")
//  }
// This kind of function is useful for code that does stuff without needing
// to return a value.
//
// Example of a function that returns an int:
// 	func functionThatReturnsAnInt() int {
//     return 42
//  }
// This kind of function is useful for code that maybe does stuff, but must
// return a value.
//
// The purpose of a function is to encapsulate some functionality in a named
// unit so that it can be reused by calling the function.
// 
// Functions can also take arguments which are values that the function uses
// to do it's work.
//
// Example of a function that takes parameters:
//  func mulitiply(x int, y int) int {
//     return x * y
//  }
// This function takes two parameters and returns the product of them.
//
// Technically, all code in Go, runs in a function. The starting point for 
// every program is the special function `main`. This function has statements
// that are executed and it can also make calls to other functions that you
// have defined.
//
// Example of a main function that calls other functions:
//  func main() {
//  	functionThatReturnsNil()
//  	number := functionThatReturnsAnInt()
//  	result := mulitiply(number, number)
//  	fmt.Println(number, "squared is", result)
//  }
// This program prints the following and then exits:
// This is what this function does. I just prints something.
// 42 squared is 1764
//


func main() {
	fmt.Println("You rolled:", roll())
}

