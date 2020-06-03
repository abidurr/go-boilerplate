package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	name, city string
	age        int
}

func (person Person) greet() string {
	return "Hello, " + person.name + "! " + "I know you are " + strconv.Itoa(person.age) + " years old."
}

func (person *Person) growOlder() {
	person.age++
}

func main() {

	person1 := Person{name: "Abidur", city: "NYC", age: 23}
	person2 := Person{"Bob", "Boston", 28}

	fmt.Println(person1, person2)
	fmt.Println(person2.city)

	fmt.Println(person2.greet())

	person2.growOlder()
	fmt.Println(person2.greet())

}
