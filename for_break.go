/* break 在for循环中跳出最内层循环，在switch和select语句中跳出整个代码块 */
package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 10; j++ {
			if j > 5 {
				break
			}
			fmt.Print(j)
		}
		fmt.Print("  ")
	}
}
