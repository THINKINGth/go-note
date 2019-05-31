package chain

func Chain(ch chan int) chan int{
	out := make(chan int)
	go func() {
		for i := range ch {
			out <- i + 1
		}
		close(out)
	}()
	return out
}