package main

import "fmt"

func main() {
	fib := fibonacci()         //这个fibonacci函数没有输入值,一定要放在循环体外面
	for i := 0; i <= 10; i++ { //在外部控制执行次数
		fmt.Println(fib()) //打印执行一次fib函数的结果

	}
}

func fibonacci() func() int {
	a, b := 0, 1
	return func() int { //返回的函数本身不需要输入值，计数是由外部执行次数决定的
		result := a
		a, b = b, a+b //移动位置的操作
		return result
	}
}
