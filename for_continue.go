/* continue 只被用于for循环中， continue 忽略剩余的循环体而直接进入下一次循环，但不是无条件地执行下一次循环，执行之前依然要满足循环的判断条件 */
package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Print(i)
		fmt.Print("  ")
	}
}
