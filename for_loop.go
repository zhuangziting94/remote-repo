package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Printf("The counter is at %d\n", i)
	}

	k := 0
START:
	fmt.Printf("The counter is at %d\n", k)
	k++
	if k < 5 {
		goto START
	}

}
