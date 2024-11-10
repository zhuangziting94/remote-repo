/*当一个函数在其函数体内调用自身，则称之为递归,用递归函数 求斐波那契数列*/
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	for n := 0; n < 41; n++ {
		x := fibonacci(n)
		fmt.Printf("fibonacci(%d) is %d\n", n, x)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}

func fibonacci(i int) (result int) {
	if i <= 1 {
		result = 1
	} else {
		result = fibonacci(i-1) + fibonacci(i-2)
	}
	return
}
