// 测试多返回值的函数的错误，var err error ,如果成功，err = nil//
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var orig string = "123" //或"ABC"
	var newS string
	fmt.Printf("The size of ints is %d\n", strconv.IntSize)
	an, err := strconv.Atoi(orig)
	if err != nil {
		fmt.Printf("orig \"%s\" is not an integer - existing with error", orig)
		return
	}
	fmt.Printf("The integer is %d\n", an)
	an = an + 5
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is %s\n", newS)
}
