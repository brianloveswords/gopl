// ch5/squares
package main

import "fmt"

// squares return a function that returns the next square number each
// time it is called
func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}
func main() {
	f1 := squares()
	f2 := squares()

	fmt.Println(f1()) // "1"
	fmt.Println(f1()) // "4"
	fmt.Println(f1()) // "9"

	fmt.Println(f2()) // "1"
	fmt.Println(f2()) // "4"
	fmt.Println(f2()) // "9"

	fmt.Println(f1()) // "16"
	fmt.Println(f2()) // "16"
}
