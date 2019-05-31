package chan_test

import (
	"fmt"
	"sync"
)


func ChanData(ch chan int) {
	for {
		select {
		    case ch <- 1:
		    case ch <- 0:
		}
		i := <- ch
		fmt.Println(i)
	}
}

func SendData(ch chan string) {
	defer close(ch)
	ch <- "北京"
	ch <- "上海"
	ch <- "宣城"
	ch <- "六盘水"
	ch <- "东京"
	fmt.Println(ch, "SendData")
}

/* 糟糕的代码
func GetData(ch, sh chan string) {
	for {
		outputS, okS := <- ch
		outputC, okC := <- sh
		if okS{
			fmt.Println(outputS)
		} else if okC{
			fmt.Println(outputC)
		} else if !okC && !okS{
			break
		}

	}
	fmt.Println(ch, sh, "GetData")
}
*/

func GetData(ch chan string, wg *sync.WaitGroup) {
	/*
	for {
		if output, ok := <- ch; ok{
			fmt.Println(output)
		}else{
			break
		}
	}
	*/
	for i := range ch{
		fmt.Println(i)
	}
	fmt.Println(ch, "GetData")
	wg.Done()
}