# 管道 #
1. **功能**：协程通信，并发控制
2. **声明**：`var ch chan int`说明存储类型，用make初始化，参数:channel类型，缓冲大小（可选），管道执行完以后必须关闭
```Go
func main() {
    strCh := make(chan string, 1)//缓冲为1
    //do something
    close(strCh) //用defer关闭较好
}
```
4. **读写**：`ch <-` 对一个管道写入数据，`<- ch` 从一个管道读取数据
```Go
func main() {
    // 如果没有缓冲区则会导致死锁
	intCh := make(chan int, 1)
	defer close(intCh)
    // 写入数据
	intCh <- 114514
    // 读取数据
	fmt.Println(<-intCh)
}
```
读取操作的第二个返回值是bool类型，表示数据是否读取成功`ints, ok := <-intCh`
5. **无缓冲**：因为无缓冲管道无法存放数据，在向管道写入数据时必须立刻有其他协程来读取数据，否则就会阻塞等待
```Go
func main() {
    ch := make(chan int)
    defer close(ch)
    //立即开启一个协程写入数据
    go func() {
      ch <- 123
    }()//go 后面可以开启函数、匿名函数、直接一个语句操作，但是不能开启内置函数
    n := <-ch //读取数据，这里如果没有缓冲区和go协程的话，需要立即用fmt.Prinln(<-ch)来接收数据
    fmt.Println(n)
}
```
6. **有缓冲**：就算有缓冲但是不用协程，同步读写很危险，一旦管道缓冲区空了或者满了，将会永远阻塞下去，因为没有其他协程来向管道中写入或读取数据。
```Go
func main() {
	// 创建有缓冲管道
	ch := make(chan int, 5)
	// 创建两个无缓冲管道
	chW := make(chan struct{})
	chR := make(chan struct{})
	defer func() {
		close(ch)
		close(chW)
		close(chR)
	}()
	// 负责写
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
			fmt.Println("写入", i)
		}
		chW <- struct{}{}
	}()
	// 负责读
	go func() {
		for i := 0; i < 10; i++ {
            // 每次读取数据都需要花费1毫秒
			time.Sleep(time.Millisecond)
			fmt.Println("读取", <-ch)
		}
		chR <- struct{}{}
	}()
	fmt.Println("写入完毕", <-chW)
	fmt.Println("读取完毕", <-chR)
}
```
这里总共创建了3个管道，一个有缓冲管道用于协程间通信，两个无缓冲管道用于同步父子协程的执行顺序。负责读的协程每次读取之前都会等待1毫秒，负责写的协程一口气做多也只能写入5个数据，因为管道缓冲区最大只有5，在没有协程来读取之前，只能阻塞等待。所以该示例输出如下
```Go
写入 0
写入 1
写入 2
写入 3
写入 4 // 一下写了5个，缓冲区满了，等其他协程来读
读取 0
写入 5 // 读一个，写一个
读取 1
写入 6
读取 2
写入 7
读取 3
写入 8
写入 9
读取 4
写入完毕 {} // 所有的数据都发送完毕，写协程执行完毕
读取 5
读取 6
读取 7
读取 8
读取 9
读取完毕 {} // 所有的数据都读完了，读协程执行完毕
```
可以看到负责写的协程刚开始就一口气发送了5个数据，缓冲区满了以后就开始阻塞等待读协程来读取，后面就是每当读协程1毫秒读取一个数据，缓冲区有空位了，写协程就写入一个数据，直到所有数据发送完毕，写协程执行结束，随后当读协程将缓冲区所有数据读取完毕后，读协程也执行结束，最后主协程退出。  

**管道中的数据流动方式与队列一样，即先进先出（FIFO），协程对于管道的操作是同步的，在某一个时刻，只有一个协程能够对其写入数据，同时也只有一个协程能够读取管道中的数据。**  

**我的尝试1**：去掉读取协程的等待时间，让读取协程与写入协程竞相进行，缓冲区有位置了才能写入，而且进程顺序不确定
```
写入 0
写入 1
写入 2
写入 3
写入 4
写入 5
读取 0
读取 1
读取 2
读取 3
读取 4
读取 5
读取 6
写入 6
写入 7
写入 8
写入 9
读取 7
读取 8
读取 9
写入完毕 {}
读取完毕 {}
```

## 利用channel优雅地让所有子协程执行完毕后再执行主协程 ##
**原理：管道的阻塞条件——阻塞等待读取**
```Go
func main() {
    // 创建一个无缓冲管道
    ch := make(chan struct{})
    defer close(ch)
    go func() {
        fmt.Println(2)
        //写入
        ch <- struct{}{}
      }()//这是匿名函数，别忘了输入参数的括号，尽管不需要输入
      //阻塞等待读取
      <- ch
      fmt.Println(1)
}
```
**用channel实现同步**
```Go
package main

import "fmt"

func printSpring(str string) { //这里是因为代码里有fmt.Print所以才不写输出吗？
	for _, data := range str {
		fmt.Printf("%c", data)
	}
	fmt.Printf("\n")
}

var ch = make(chan int)
var tongBu = make(chan int)

func person1() {
	printSpring("Gerald")
	tongBu <- 1
	ch <- 1
} //函数套函数

func person2() {
	<-tongBu //tongBu阻塞等待读取，确保两协程的同步，不会乱顺序
	printSpring("Seligman")
	ch <- 2
}

func main() {
	//目的：使用channel来实现person1先于person2执行
	go person1()
	go person2()
	count := 2
	//判断所有协程是否退出
	for range ch {
		count--
		if 0 == count {
			close(ch)
		}
	}
	//再在主协程里做其他事，do someting
}
```
* count 表示有所少个协程
* ch 用来子协程与主协程之间的同步
* tongBu 用来两个协程之间的同步
* 主协程阻塞等待数据，每当一个子协程执行完后，就会往 ch 里面写一个数据，主协程收到后会使 count–，当 count 减为 0，关闭 ch，主协程将不阻塞在 range ch。
