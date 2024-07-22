package main

import "fmt"

func main() {
	a := 3
	b := 4

	fmt.Println(Add(a, b))
}

func Add(a, b int) int {
	return a + b
}