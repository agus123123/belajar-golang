package main

import (
	"fmt"
)

func updateName(x *string) string {
	*x = "Gita"
	return *x
}
func main() {
	name := "Agus"

	// fmt.Println("Before update:", &name)
	m := &name
	// fmt.Println("After update:", m)
	fmt.Println(name)

	// fmt.Println("value of m:", *m)

	updateName(m)
	fmt.Println(name)

}
