/*
如果函数的最后一个参数是采用 ...type 的形式，那么这个函数就可以处理一个变长的参数，这个长度可以为 0，这样的函数称为变参函数。
如果参数被存储在一个数组 arr 中，则可以通过 arr... 的形式来传递参数调用变参函数。
*/
package main

import "fmt"

func Min(arr ...int) int {
	if len(arr) == 0 {
		return 0
	}
	min := arr[0]
	for _, val := range arr {
		if val < min {
			min = val
		}
	}
	return min
}

func main() {
	x := Min(87, 5, 6, 3)
	fmt.Printf("The minimum is %d\n", x)
	arr := []int{7, 5, 9, 1, 3}
	y := Min(arr...)
	fmt.Printf("The minimum in the array arr is %d\n", y)
}
