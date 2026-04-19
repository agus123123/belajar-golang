package main

import (
	"fmt"
)

func updateName(x string) string {
	x = "Gita"
	return x
}

func updateMenu(y map[string]float64) {
	y["naspad"] = 25000
}
func main() {
	name := "Agus"

	name = updateName(name)
	fmt.Println(name)

	menu := map[string]float64{
		"naspad": 20000,
		"laptop": 5000,
		"eskrim": 10000,
	}

	updateMenu(menu)
	fmt.Println(menu)

}
