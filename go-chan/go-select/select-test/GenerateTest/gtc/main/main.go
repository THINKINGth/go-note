package main

import (
	"go-note/go-chan/go-select/select-test/GenerateTest/gtc"
	"fmt"
	"runtime"
	"time"
)

func main() {

	does := make(chan struct{})
	ch := gtc.GenerateT(does)
	for i := 0; i < 2; i++ {
		fmt.Println(<-ch)
	}
	fmt.Println(runtime.NumGoroutine())
	does <- struct{}{}
	fmt.Println(runtime.NumGoroutine())
	for i := range ch{
		fmt.Println(i)
	}
	time.Sleep(5 * time.Second)
    fmt.Println(runtime.NumGoroutine())
}