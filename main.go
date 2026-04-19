package main

import (
	"fmt"
	"math"
)

func sayHello(n string) {
	fmt.Printf("Halo selamat %v \n", n)
}
func sayBye(n string) {
	fmt.Printf("Sampai jumpa %v \n", n)
}
func cyclyeNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}
func main() {
	// sayHello("Alice")
	// sayBye("Alice")

	cyclyeNames([]string{"cloud", "tifa", "agus"}, sayHello)
	cyclyeNames([]string{"cloud", "tifa", "agus"}, sayBye)

	a1 := circleArea(10.5)
	a2 := circleArea(12.5)

	fmt.Println(a1, a2)
	fmt.Printf("Luas lingkaran dengan jari-jari 10.5 adalah %v \n", a1)
	fmt.Printf("Luas lingkaran dengan jari-jari 12.5 adalah %v \n", a2)
}
