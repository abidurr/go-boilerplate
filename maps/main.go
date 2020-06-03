package main

import "fmt"

func main() {

	ids := map[int]string{1: "Alice", 2: "Bob", 3: "Cath"}
	
	fmt.Println(ids)

	emails := make(map[string]string)

	emails["Alice"] = "alice@email.com"
	emails["Bob"] = "bob@email.com"
	emails["Cath"] = "cath@email.com"

	fmt.Println(emails)
	
	delete(emails, "Cath")
	fmt.Println(emails)

}