package main

import (
	"go-note/go-chan/go-select/select-test/GenerateTest/gtb/funcgtb"
	"fmt"
	"runtime"
)

func main() {
	does := make(chan struct{})
	fmt.Println(runtime.NumGoroutine())
	ch := funcgtb.GenerateT(does)
	fmt.Println(runtime.NumGoroutine())
	fmt.Println(<-ch)
	fmt.Println(runtime.NumGoroutine())
	close(does)
	//fmt.Println("close-main:", len(ch), cap(ch))
	fmt.Println(runtime.NumGoroutine())
	for i := range ch {
		//fmt.Println("main-for:", runtime.NumGoroutine(),len(ch), cap(ch))
		fmt.Println(i)
	}
	fmt.Println(runtime.NumGoroutine())

}