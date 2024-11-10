在 Go 中，byte 类型切片（[]byte）是一种常用的数据结构，特别适合处理字节数据，比如字符串、文件数据、网络数据等。byte 本质上是 uint8 的别名，用于表示一个字节。下面是定义和使用 byte 切片的几种方法：
# 定义 []byte 类型的切片 #
## 直接定义并初始化 ##
```Go
data := []byte{65, 66, 67} // 定义一个包含字节 65, 66, 67 的切片
fmt.Println(data)           // 输出 [65 66 67]
fmt.Println(string(data))    // 输出 "ABC"
```
## 将字符串转换为 []byte ##
```Go
str := "Hello"
data := []byte(str)
fmt.Println(data)           // 输出 [72 101 108 108 111]
fmt.Println(string(data))    // 输出 "Hello"
```
## 使用 make 函数创建指定长度的 []byte
可以使用 make 函数创建一个指定长度和容量的 []byte 切片。
```Go
data := make([]byte, 5)    // 创建长度和容量都为 5 的切片，默认值为 0
fmt.Println(data)          // 输出 [0 0 0 0 0]
```
