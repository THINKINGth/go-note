package main

import (
	"go-note/go-chan/test-add/test-Task"
	"sync"
	"fmt"
)

func main() {

	num := 50000 // 利用goroutine求[0 ,num]范围内整数的和。
	sc := 1000  // 划分的组数
	chanTask := make(chan test_Task.Task, sc)
	result := make(chan int, sc)
	var wg sync.WaitGroup
	var sg sync.WaitGroup
	sg.Add(1)
	go test_Task.InitTask(chanTask, result, num + 1, sc)
	go test_Task.ProcessTask(chanTask, result, &wg, &sg)
	sum := test_Task.ProcessResult(result, &wg, &sg)
	fmt.Println(sum)
	sg.Wait()

}
