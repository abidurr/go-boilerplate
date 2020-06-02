package main

import "fmt"

func main() {

	charArray := [2]string{"A", "B"}
	fmt.Println(charArray)

	charSlice := []string{"C", "D", "E", "F"}
	fmt.Println(charSlice)

	fmt.Println(len(charSlice))
	fmt.Println(charSlice[1:2])

}