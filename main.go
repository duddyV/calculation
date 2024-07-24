// Package main provides a simple example of adding two integers using the Add function.
package main

import "fmt"

// Add takes two integers a and b and returns their sum.
// For more information on addition, see https://www.mathsisfun.com/numbers/addition.html.
func Add(a, b int) int {
	return a + b
}

func main() {
	a := 3
	b := 4

	fmt.Println(Add(a, b))
}
