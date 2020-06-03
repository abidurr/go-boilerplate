package main

import "fmt"

func main() {

	ids := []int{23, 34, 54, 67, 78, 90}

	for i, id := range ids {
		fmt.Printf("ID#%d: %d\n", i, id)
	}

	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum is", sum)

	emails := map[string]string{"Alice": "alice@email.com", "Bob": "bob@email.com"}
	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}
}