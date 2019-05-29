package main

import (
	"fmt"
	"runtime"
	"time"
	"go-note/go-chan/go-channel/core-channel"
)

// "不要通过共享内存来通信，而是通过通信来共享内存"
// 通道是Go通过通信来实现共享内存的载体
// 声明： var variate chan datatype
// 赋值： variate = make(chan datatype, bufferNum)
/*
 Get
 1. 什么是无缓冲的通道和有缓冲的通道
 2. 有缓冲通道中,len() , cap()的意义（同切片进行比较）
 3. 无缓冲通道， 有缓冲通道的使用
 4. c <- struct{}{}
 5. 验证了golang的块级作用域的效果
 6. 匿名结构体（同匿名函数进行比较）
*/

func main() {
	Ch := make(chan int)
	ChanBuffer := make(chan int, 15)
	fmt.Println("buffer: ", len(ChanBuffer), " No buffer: ", len(Ch)) // 没有被读取的元素数
	fmt.Println("buffer: ", cap(ChanBuffer), " No buffer: ", cap(Ch)) // 代表整个通道的容量
	//  无缓冲通道主要用于同步和通信， 有缓冲通道主要用于通信
	//  goroutine运行后退出， 写到缓冲通道中的数据不会消失, !!它可以缓冲和适配两个goroutine处理速率不一样的问题!!
	//  缓冲通道和消息队列类似，有消峰和增大吞吐量的功能
	c := make(chan struct{})
	go core_channel.TestChanRead(ChanBuffer, c)
	// 此时ChanBuffer已经关闭，并且goroutine执行的函数TestChanRead已经结束
	fmt.Println("此时的goroutine数： ", runtime.NumGoroutine())
	<- c // 读通道，通过通道进行同步等待
	fmt.Println("buffer: ", len(ChanBuffer), " No buffer: ", len(Ch)) // 没有被读取的元素数
	fmt.Println("buffer: ", cap(ChanBuffer), " No buffer: ", cap(Ch)) // 代表整个通道的容量
	//  虽然通道已经关闭，但还可以继续从缓冲区里读取数据，但读取过后将不再会有数据
	//  利用for variate := chanVariable{} 时，一定要保证通道已经关闭，不然会造成阻塞。
	for i := range ChanBuffer{
		fmt.Printf("%d ", i)
	}
	{
		for i := range ChanBuffer{
			fmt.Println("没有执行")
			fmt.Printf("%d ", i)
		}
		j ,ok := <- ChanBuffer // 如果通道没有关闭且通道中没有数据，读取该通道会导致goroutine阻塞
		if !ok{
			fmt.Println("\n可以读取已经关闭的通道")
			fmt.Println(j)
		}
	}

	fmt.Println("buffer: ", len(ChanBuffer), " No buffer: ", len(Ch)) // 没有被读取的元素数
	fmt.Println("buffer: ", cap(ChanBuffer), " No buffer: ", cap(Ch)) // 代表整个通道的容量

    // 无缓冲的通道
	go func() {
		Ch <- 1
	}()
	time.Sleep(2 * time.Second)
	fmt.Println("buffer: ", len(ChanBuffer), " No buffer: ", len(Ch)) // 没有被读取的元素数
	fmt.Println("buffer: ", cap(ChanBuffer), " No buffer: ", cap(Ch)) // 代表整个通道的容量
}