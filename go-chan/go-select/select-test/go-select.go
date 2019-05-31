package select_test

import (
	"time"
	"fmt"
	"reflect"
)

func TestSelect(ch chan int) int{
	var i int
	ins := make(chan bool)
	go func (){
		time.Sleep(3e9)
		ins <-true
	}()
	fmt.Println(ch, reflect.TypeOf(ch), reflect.ValueOf(ch))
	select {
	    case <- ch:
		    fmt.Println("数据传输成功")
	    case <- ins:
		    fmt.Println("数据传输失败")
	}
	return i
}