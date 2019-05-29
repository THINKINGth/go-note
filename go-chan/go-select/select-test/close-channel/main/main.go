package main

import (
	"go-note/go-chan/go-select/select-test/close-channel"
	"fmt"
	"runtime"
)

func main () {
	done := make(chan struct{})
	ch := close_channel.GenerateInt(done)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	close(done)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(runtime.NumGoroutine())
}