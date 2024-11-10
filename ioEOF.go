package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	reader := strings.NewReader("Hello, Go!")
	buf := make([]byte, 4)

	for {
		n, err := reader.Read(buf)
		if err == io.EOF {
			fmt.Println("Reach the end of the file")
			break
		}

		if err != nil {
			fmt.Println("Error reading :", err)
			return
		}
		fmt.Printf("Read %d bytes: %s\n", n, buf[:n])
	}
}

//io.EOF 常用于读取文件、网络数据流等操作中，表示已经读取完所有的数据。这在处理循环读取操作时特别重要，因为可以通过检查是否遇到 io.EOF 来决定何时停止读取。
