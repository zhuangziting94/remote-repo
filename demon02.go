package main

import "fmt"

func main() {
	p := Add2()
	fmt.Println(p(1))
	fmt.Println(p(2))
	fmt.Println(p(3))
	fmt.Println(p(4))
}

func Add2() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}
