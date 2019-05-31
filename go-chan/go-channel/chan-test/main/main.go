package main

import (
	"go-note/go-chan/go-channel/chan-test"
	"sync"
)

func main () {
	/* go chan_test.ChanData
	var ch chan int
	ch = make(chan int, 1)
	go chan_test.ChanData(ch)
	time.Sleep(1e9)
	*/
	var ch chan string
	ch = make(chan string)
	sh := make(chan string)
	var wg sync.WaitGroup
	go chan_test.SendData(ch)
	go chan_test.SendData(sh)
	wg.Add(1)
	go chan_test.GetData(ch, &wg)
	wg.Add(1)
	go chan_test.GetData(sh, &wg) // 小心主进程结束终止其他协程
	wg.Wait()

}