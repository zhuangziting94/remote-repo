package main

import "fmt"

func main() {
	s := "good bye"
	fmt.Printf("Here is the initial string s: %s\n", s)
	var p *string
	p = &s
	*p = "good night"
	fmt.Printf("Here is the printer: %p\n", p)
	fmt.Printf("Here is the string *p: %s\n", *p)
	fmt.Printf("Here is the changed string s: %s\n", s)
}
