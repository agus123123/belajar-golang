package main

import (
	"fmt"
)

func main() {
	menu := map[string]float64{
		"Burger": 5.99,
		"Pizza":  8.99,
		"Salad":  4.99,
		"Soda":   1.99,
	}

	fmt.Println(menu)
	fmt.Println(menu["Pizza"])

	// loping

	for k, v := range menu {
		fmt.Println(k, ":", v)
	}

	//ints as key type

	phonebook := map[int]string{
		1234567890: "Alice",
		9876543210: "Bob",
		5555555555: "Charlie",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[1234567890])

	phonebook[1234567890] = "David"
	fmt.Println(phonebook)

	phonebook[5555555555] = "Agus"
	fmt.Println(phonebook)
}
