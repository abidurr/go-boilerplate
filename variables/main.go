// := is block scoped, var and const are global scoped

package main

import "fmt"

func main() {

	const name string = "Abidur"
	var age int32 = 23
	coding := true

	fmt.Println(name, age, coding)

	fmt.Printf("%T", coding)

}