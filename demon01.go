package main

import "fmt"

//函数功能：求和
//函数名字： getSum,参数为空
//getSum函数返回值为一个函数，这个函数的参数是一个int 类型的参数，返回值也是int类型
func getSum() func(int) int {
	var sum int = 0
	return func(num int) int {
		sum = sum + num
		return sum
	}
}

//闭包 ： 返回的匿名函数+匿名函数以外的变量sum
//感受：匿名函数中引用的变量会一直保存在内存中，可以一直使用，达到累加效果
func main() {
	f := getSum()
	fmt.Println(f(1)) //1,此处sum传回来后彻底变成了1，故下面是1+2=3
	fmt.Println(f(2)) //3
	fmt.Println(f(3)) //6
}
