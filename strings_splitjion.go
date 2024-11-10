/*
strings.Fields(s) 将会利用 1 个或多个空白符号来作为动态长度的分隔符将字符串分割成若干小块，并返回一个 slice，如果字符串只包含空白符号，则返回一个长度为 0 的 slice。

strings.Split(s, sep) 用于自定义分割符号来对指定字符串进行分割，同样返回 slice。

因为这 2 个函数都会返回 slice，所以习惯使用 for-range 循环来对其进行处理
Join 用于将元素类型为 string 的 slice 使用分割符号来拼接组成一个字符串
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "The quick brown fox jumps over the lazy dog"
	sl := strings.Fields(str)
	fmt.Printf("Splitted in slice: %v\n", sl)
	for _, val := range sl {
		fmt.Printf("%s-", val)
	}
	fmt.Println()
	str2 := "GO1|The ABC of Go|25"
	sl2 := strings.Split(str2, "|")
	for _, val := range sl2 {
		fmt.Printf("%s-", val)
	}
	fmt.Println()
	str3 := strings.Join(sl2, ";")
	str4 := strings.Join(sl, ";")
	fmt.Printf("sl2 joined by ;: %s\n", str3)
	fmt.Printf("sl1 joined by ;: %s\n", str4)
}
