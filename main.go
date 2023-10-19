package main

import "fmt"

func sum(a, b int) int {
	return a + b
}

func div(a, b int) int {
	return a / b
}

func display(msg string) {
	fmt.Println(msg)
}
func main() {
	fmt.Println("Hello, Github")
	fmt.Println(sum(1, 2))
	fmt.Println(div(1, 2))
	display("Hello GitFlow from Me")
}
