package test_Task

import (
	"sync"
)

type Task struct{
	start int
	end int
	result chan int
}

func do(task Task, wg *sync.WaitGroup){
	sum := 0
	for i := task.start; i < task.end; i++ {
		sum += i
	}
	task.result <- sum
    wg.Done()
}

func InitTask(chanTask chan Task, resultTask chan int, num int, sc int) {
	p := num / sc  // 组内元素数
	mod := num - (sc * p) // 剩余的元素自成一组
	e := p * sc
	for i := 0; i < sc; i++ {
		task := Task{
			start: i * p,
			end: (i + 1) * p,
			result: resultTask,
		}
		// fmt.Println(task)
		chanTask <- task
	}
	// 有剩余元素
	if mod != 0 {
		task := Task{
			start: e,
			end : e + mod,
			result: resultTask,
		}
		//fmt.Println(task)
		chanTask <- task
	}
	close(chanTask)
}


func ProcessTask(chanTask chan Task, result chan int, wg *sync.WaitGroup, sg *sync.WaitGroup) {
	for i := range chanTask{
		wg.Add(1)
		go do(i, wg)
	}
	wg.Wait()
	close(result)
	// 保证执行ProcessResult函数的时候，goroutine 所执行的do()已经全部执行完毕
}

func ProcessResult(result chan int, wg *sync.WaitGroup, sg *sync.WaitGroup) int{
	sum := 0
	for i := range result{
		// 证明在通道关闭之前，数据是一个接着一个送过来的。
		// fmt.Println(len(result), cap(result))
		sum += i
	}
	sg.Done()
	return sum
}



/*
package test_Task

import (
	"sync"
	"fmt"
)

type Task struct{
	start int
	end int
	result chan int
}

func do(task Task, otherTask chan Task, wg *sync.WaitGroup){
	sum := 0
	for i := task.start; i < task.end; i++ {
		sum += i
	}
	task.result <- sum
    otherTask <- task
    wg.Done()
}

func InitTask(chanTask chan Task, num int, sc int, wg *sync.WaitGroup) {
	p := num / sc; // 分组数
	mod := num % sc; // 剩余的元素自成一组
	e := p * sc
	for i := 0; i < p; i++ {
		task := Task{
			start: i * sc,
			end: (i + 1) * sc,
			result: make(chan int, 1),
		}
		//fmt.Println(task)
		chanTask <- task
	}
	// 有剩余元素
	if mod != 0 {
		task := Task{
			start: e,
			end : e + mod,
			result: make(chan int, 1),
		}
		chanTask <- task
	}
	close(chanTask)
}


func ProcessTask(chanTask chan Task, result chan int, wg *sync.WaitGroup, sg *sync.WaitGroup) {
	ch := make(chan Task, 10)
	for i := range chanTask{
		wg.Add(1)
		go do(i, ch, wg)
	}
	wg.Wait()
	go ProcessResult(ch, result, wg, sg)
	// 保证执行ProcessResult函数的时候，goroutine 所执行的do()已经全部执行完毕
}

func ProcessResult(chanTask chan Task, result chan int, wg *sync.WaitGroup, sg *sync.WaitGroup) {
	sum := 0
	close(chanTask)
	for i := range chanTask {
		sum += <- i.result
	}
	result <- sum
	fmt.Println(<-result)
	sg.Done()
}

 */