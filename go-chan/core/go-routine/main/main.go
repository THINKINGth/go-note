package main

import (
	"fmt"
	"go-note/go-chan/core/go-routine/func-routine"
	"runtime"
	"time"
)

/*
 Get
 1. 验证了使用goroutine的函数没有返回值
 2. 当前版本的golang可以智能的检测核心数（当前把版本1.10.3）
 3. 验证了runtime.Goexit()的功能
 4. 验证了runtime.Gosched()的功能
 5. time.sleep()的新用法
 Error
 1. for循环
 2. make(chan int)
 3  仔细检查通道信息，避免主进程死锁
*/
func main(){
	num := 100000
	//runtime.GOMAXPROCS(1) // 确认，当前版本的golang已经实现了智能检测核心数
	go func (num int) {
		var sum int
		for i := 0; i < num; i++ {
			sum += i
		}
		fmt.Println(sum, "func()")
	}(num)
	ch := make(chan int)
	go func_routine.Counts(ch, num)
	fmt.Println("运行的goroutine数: ", runtime.NumGoroutine())
	SumFunc := <- ch
	fmt.Println(SumFunc)
	time.Sleep(5 * time.Second)
}
