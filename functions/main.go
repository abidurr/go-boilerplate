package main

import "fmt"

func greeting(name string) string {
	return "Hello from the greeting function, " + name
}

func getSum(num1, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println(greeting("stranger"))
	fmt.Println(getSum(5, 6))
}