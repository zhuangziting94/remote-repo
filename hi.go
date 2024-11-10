package main

import (
	"fmt"
	"strconv"
)

func atoi(s string) (n int) {
	n, _ = strconv.Atoi(s)
	return
}

func main() {
	var str = "123"
	newI := atoi(str)
	fmt.Println(newI)
}
