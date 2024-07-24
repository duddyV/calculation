// Package main provides a simple example of adding two integers using the Add function.
package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

// Add takes two integers a and b and returns their sum.
// For more information on addition, see https://www.mathsisfun.com/numbers/addition.html.
func Add[T Number] (a, b T) T {
	return a + b
}

func main() {
	a := 33.1
	b := 4.37
	c := 4
	d := 111

	fmt.Println(Add(a, b))
	fmt.Println(Add(c, d))
}
