package main

import "fmt"

func main() {

	x := 3
	y := 7
	if x < y {
		fmt.Printf("%d < %d", x, y)
	} else if x == y {
		fmt.Printf("%d = %d", x, y)
	} else {
		fmt.Printf("%d > %d", x, y)
	}

	color := "blue"
	switch color {
	case "red":
		fmt.Printf("\nred")
	case "blue":
		fmt.Printf("\nblue")
	case "green":
		fmt.Printf("\ngreen")
	default:
		fmt.Printf("not primary color")

	}
}