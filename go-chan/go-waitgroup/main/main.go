package main

import (
	"sync"
	"net/http"
	"fmt"
)

/**
  1. 熟悉使用godoc
  2. var, type
  3. 字符切片的使用范例(多关注数据结构的使用)
  4. 掌握WaitGroup()在进程间通信的基本使用方法
  5. 熟悉错误处理的范式
 */

var wg sync.WaitGroup
var urls = []string{
	"http://www.sina.com/",
	"http://www.bilibili.com/",
	"http://www.baidu.com/",
}
func main() {
	for _, url := range urls {
		// Increment the WaitGroup counter.
		wg.Add(1)
		// Launch a goroutine to fetch the URL.
		go func(url string) {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()
			// Fetch the URL.
			resp, err := http.Get(url)
			if err == nil {
				fmt.Println("目标url: ", url)
				fmt.Println(resp.Status)
			}
		}(url)
	}
	// Wait for all HTTP fetches to complete.
	wg.Wait()
}
