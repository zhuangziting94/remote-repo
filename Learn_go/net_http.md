# 访问网址 #
- GET:
```Go
package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// 定义目标网址
	url := "https://example.com"

	// 发送 HTTP GET 请求
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("请求失败:", err)
		return
	}
	defer resp.Body.Close() // 确保在函数结束时关闭响应体

	// 检查 HTTP 状态码
	if resp.StatusCode != http.StatusOK {
		fmt.Println("HTTP 请求失败，状态码:", resp.StatusCode)
		return
	}

	// 读取响应内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应内容失败:", err)
		return
	}

	// 打印响应内容
	fmt.Println(string(body))
}
```
http.Get:

用于发送一个 HTTP GET 请求。
返回一个 *http.Response 对象和错误信息。
关闭响应体:

defer resp.Body.Close() 确保响应体资源被正确释放，避免内存泄漏。
检查状态码:

检查 resp.StatusCode 是否为 http.StatusOK (200) 确保请求成功。
读取响应内容:

使用 io.ReadAll 读取响应体的内容。
运行结果
运行该程序后，会将目标网址的 HTML 或响应内容打印到控制台。  

- 如果需要更复杂的功能（例如 POST 请求、设置头部、超时控制等），可以使用 http.Client：
```Go
package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: 10 * time.Second, // 设置超时时间
	}

	req, err := http.NewRequest("GET", "https://example.com", nil)
	if err != nil {
		fmt.Println("创建请求失败:", err)
		return
	}

	// 添加请求头
	req.Header.Set("User-Agent", "MyCustomUserAgent/1.0")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("请求失败:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应内容失败:", err)
		return
	}

	fmt.Println(string(body))
}
```
简单访问: 使用 http.Get 即可满足需求。
复杂需求: 使用 http.Client 和 http.NewRequest 自定义请求。
