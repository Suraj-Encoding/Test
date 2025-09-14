package main

import "fmt"

func Square(a int) int {
	return a * a
}

func Cube(a int) int {
	return a * a * a
}

func main() {
	fmt.Println("Square: ", Square(2))
	fmt.Println("Cube: ", Cube(2))
}
