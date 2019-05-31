package close_channel

import (
	"math/rand"
)

func GenerateInt(done chan struct{}) (ch chan int) {
	ch = make(chan int)
	go func () {
		Lab:
			for {
			    select {
			    // Int returns a non-negative pseudo-random int from the default Source.
			    case ch <- rand.Int():
			    case <- done:
			    	break Lab
			}
		}
		close(ch)
	}()
	return ch
}