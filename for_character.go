package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("G")
		}
		fmt.Println()
	}
	str := "G"
	for k := 0; k < 5; k++ {
		fmt.Print(str)
		str += "G"
		fmt.Println()
	}
}
