package main

import "fmt"

func Test(a int) int {
	return a * a
}

func main() {
	fmt.Println("Hello World")
	ans := Test(2)
	fmt.Println("answer: ", ans)
}
