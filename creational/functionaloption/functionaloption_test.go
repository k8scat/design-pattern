package functionaloption

import (
	"fmt"
	"time"
)

func ExampleFunctionalOption() {
	// 使用自定义超时时间创建一个新的HTTP客户端
	client := NewHTTPClient(WithTimeout(5 * time.Second))

	// 使用自定义HTTP客户端发起GET请求
	resp, err := client.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Status Code:", resp.Status)
	// Output:
	// Status Code: 200 OK
}
