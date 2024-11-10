package main

import "fmt"

func main() {
	var i1 = 5
	fmt.Printf("An integer: %d, it's location in memory is %p.\n", i1, &i1)

	var intP *int
	intP = &i1
	fmt.Printf("The value at memory location %p is %d\n", intP, *intP)
}
