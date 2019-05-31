package main

import (
	"go-note/go-chan/go-select/select-test/GenerateTest/gta/funcgta"
	"fmt"
)

func main(){
	ch := funcgta.GenerateA()
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	close(ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}