package funcgta
import (
	"math/rand"
)

/*
 func GenerateA() (ch chan int)
 (ch chan int)会声明一个int类型的通道ch
 */
func GenerateA() (ch chan int){
	ch = make(chan int)
	go func() {
		for {
			select {
			case ch <- rand.Int():
			}
		}
	}()
	return ch
}