package main

import "fmt"

// calculate the square of input values
func Square(a int) int {
	return a * a
}

func main() {
	fmt.Println("main function")
	fmt.Println(Square(2))
}