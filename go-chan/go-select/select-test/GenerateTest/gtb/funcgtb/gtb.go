package funcgtb

import (
	"math/rand"
)

func GenerateA() chan int{
	ch := make(chan int)
	go func() {
		for {
			ch <- rand.Int()
		}
	}()
	return ch
}

func GenerateB() chan int{
	ch := make(chan int)
	go func() {
		for {
			ch <- rand.Int()
		}
	}()
	return ch
}

/**
 如何协调两个通道传输数据, 并打印传输的数据
 写程序之前想一想自己写的代码目的，原理，思路，不能一叶障目对不对，自己用到的结构原理是什么？
 在扇入的情况下，如何关闭函数
 */
func GenerateT(does chan struct{}) chan int{
	ch := make(chan int, 10)
	//Gra := GenerateA()
	//Grb := GenerateB()
	go func() {
		Lab:
		for {
			select {
			case ch <- <- GenerateA():
			case ch <- <- GenerateB():
			case <- does:
				break Lab
			}
		}
		close(ch)
	}()
	return ch
}