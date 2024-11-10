## 用golang怎么能够遍历一个数组的字符，并将当前字符和前一个字符不相同的字符拷贝至另一个数组 ##
在 Go 中，可以遍历一个字符串或字符数组，将与前一个字符不相同的字符拷贝到另一个数组中。这里有一个示例，假设我们要将非重复的字符从源数组拷贝到目标数组中：
```Go
package main

import (
    "fmt"
)

func copyUniqueChars(arr []rune) []rune {
    // 如果输入数组为空，则直接返回空的结果数组
    if len(arr) == 0 {
        return []rune{}
    }

    // 初始化结果数组，将第一个字符添加到结果数组中
    result := []rune{arr[0]}

    // 遍历 arr，从第二个字符开始
    for i := 1; i < len(arr); i++ {
        // 如果当前字符和前一个字符不相同，则添加到结果数组
        if arr[i] != arr[i-1] {
            result = append(result, arr[i])
        }
    }

    return result
}

func main() {
    s := "aabbccddeeffgghhii"  // 示例字符串
    arr := []rune(s)           // 将字符串转换为字符切片（[]rune），适用于多字节字符

    uniqueChars := copyUniqueChars(arr)
    fmt.Println("Original array:", string(arr))
    fmt.Println("Unique chars array:", string(uniqueChars))
}
```
* **字符数组类型**：`[]rune` 用于处理 Unicode 字符，保证可以支持多字节字符。
* **转换为字符串输出**：在 `fmt.Println` 中使用 `string` 将 `[]rune` 转为字符串，以便于直观显示。
