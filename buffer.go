package main

import (
	"bytes"
	"fmt"
)

func main() {
	sl := bytes.NewBuffer([]byte{9, 10, 6})
	sl.Write([]byte{1, 3, 5})
	fmt.Println(sl.Bytes())
}
