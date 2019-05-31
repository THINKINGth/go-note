package core_channel

import (
	"fmt"
	"runtime"
)

// "不要通过共享内存来通信，而是通过通信来共享内存"
// 通道是Go通过通信来实现共享内存的载体
// chan datatype

/*
 Get
 1. 什么是无缓冲的通道和有缓冲的通道
 2. 有缓冲通道中,len() , cap()的意义
 */
 func TestChanRead(ch chan int, c chan struct{}) {
 	fmt.Println("此时的goroutine数： ", runtime.NumGoroutine())
 	for i := 0; i < 10; i++ {
 		ch <- i
	}
	close(ch)
	c <- struct{}{}

 }