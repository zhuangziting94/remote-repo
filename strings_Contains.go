package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "This is an example of a string"
	fmt.Printf("T/F? Does the string \"%s\" have the word %s?", str, "string")
	fmt.Printf("%t\n", strings.Contains(str, "string"))
}
