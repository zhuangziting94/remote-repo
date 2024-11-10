package main

import "fmt"

//将匿名函数赋给全局变量
var func01 = func(num1 int, num2 int) int {
	return num1 * num2
}

func main() {
	//调用匿名函数，定义的同时并调用，只用一次
	result := func(a int, b int) int {
		return a + b
	}(10, 20)

	fmt.Println(result)

	//将匿名函数 赋给一个变量(这个变量是局部变量），该变量是一个函数类型的变量
	sub := func(a int, b int) int {
		return a - b
	}
	//调用sub
	result01 := sub(70, 30)
	fmt.Println(result01)

	result02 := func01(10, 40)
	fmt.Println(result02)
}
