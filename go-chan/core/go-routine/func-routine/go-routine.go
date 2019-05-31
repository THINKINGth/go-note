package func_routine

import (
	"runtime"
	"fmt"
)

func Counts (ch chan int, num int) {
	sum := 0
	runtime.Gosched() // 让出当前的控制权
	defer func(){
		fmt.Println("goroutine提前退出", "funcCounts")
		ch <- 0
	}()
	// goroutine提前退出, 退出之前执行defer语句
	// goroutine 在结束当前运行前，会提前执行注册的defer语句
	runtime.Goexit()
	for i := 0; i < num; i ++ {
		sum += i
	}
    // 没有提前注册
    defer fmt.Println("没有提前注册")
	ch <- sum
}