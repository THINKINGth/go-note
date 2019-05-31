package gtc

import (
	"math/rand"
)

func GenerateA(cl chan struct{}) chan int{
	ch := make(chan int)
	go func(){
	Lga:
		for {
			select {
			case ch <- rand.Int():
			case <- cl:
				break Lga
			}
		}
		close(ch)
	}()
	return ch
}

func GenerateB(cl chan struct{}) chan int{
	ch := make(chan int)
    go func() {
	Lgb:
		for {
			select {
			case ch <- rand.Int():
			case <- cl:
				break Lgb
			}
		}
		close(ch)
	}()
	return ch
}

func GenerateT(does chan struct{}) chan int{
	ch := make(chan int, 20)
	send := make(chan struct{})
	go func(){
		Lab:
			for{
				select {
				case ch <- <- GenerateA(send):
				case ch <- <- GenerateB(send):
				case <- does:
					send <- struct{}{}
					send <- struct{}{}
					break Lab
				}
			}
			close(ch)
	}()
	return ch
}