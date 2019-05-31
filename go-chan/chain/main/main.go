package main

import (
	"fmt"
	"go-note/go-chan/chain/chain"
)

/*
 再一次熟悉goroutine发生阻塞的条件
 */
func main() {
	in := make(chan int)
	// 利用chain.Chain实现管道功能
	// 启动器
	go func() {
		if cap(in) == 0{
			in <- 1
		}
		for i := 0; i < cap(in); i++ {
			in <- 1
		}
		close(in)
	}()
	ch := chain.Chain((chain.Chain(chain.Chain(in))))
	for i := range ch{
		fmt.Println(i)
	}
}
