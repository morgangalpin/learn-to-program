// Author: Morgan the great teacher!
/* This is a 
	multiline 
	comment.
*/
package main

import (
	"fmt"
)

func main() {
	var number int = 5
	fmt.Println("The value of number is:", number)
	number = 2 + number
	fmt.Println("The value of number is now:", number)

	var text = "This is text"
	fmt.Println("The value of text is:", text)

	pi := 3.14
	fmt.Println("The value of pi is:", pi)
	pi = pi * 2
	fmt.Println("The value of 2 times pi is:", pi)
}
