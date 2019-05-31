package main

import (
	"go-note/go-chan/go-select/select-test"
	"fmt"
	"reflect"
)

func main() {
	var ch chan int
	ch = make(chan int)
	go func (){
		ch <- 1
	}()
	fmt.Println(ch, reflect.TypeOf(ch), reflect.ValueOf(ch))
	select_test.TestSelect(ch)

}