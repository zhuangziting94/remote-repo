package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%3 == 0 && i%5 != 0:
			fmt.Print("Fizz ")
		case i%5 == 0 && i%3 != 0:
			fmt.Print("Buzz ")
		case i%5 == 0 && i%3 == 0:
			fmt.Print("FizzBuzz ")
		default:
			fmt.Printf("%d ", i)
		}
	}
}
