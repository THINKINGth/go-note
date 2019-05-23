# go-select

> **线程**
```go 
func main() {
	var ch chan int
	ch = make(chan int)
	go func (){
		ch <- 1
	}()
	fmt.Println(ch, reflect.TypeOf(ch), reflect.ValueOf(ch))
	go_select.TestSelect(ch)
}

func TestSelect(ch chan int) {
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
}
```
第四行的匿名函数创造了一个协程，如果将第七行的语句改成：
```go
fmt.Println(<-ch, ch, reflect.TypeOf(ch), reflect.ValueOf(ch))
```
此时```<- ch```会导致出现```ch```提前接收到数据。{（主线程阻塞，等待ch接收数据）
特别的，如果通道```ch```一直没有接收到数据，就会导致出现死锁。}而函数```TestSelect```
的通道中的```ch```无法写入数据，导致数据传输失败。
```go
func main() {
	var ch chan int
	ch = make(chan int)
	func (){
		ch <- 1
	}()
	fmt.Println(ch, reflect.TypeOf(ch), reflect.ValueOf(ch))
	go_select.TestSelect(ch)
```

```
fatal error: all goroutines are asleep - deadlock!
goroutine 1 [chan send]:
main.main.func1(...)
	C:/Users/Rongqin/go/src/go-note/go-chan/main/go-select/main/main.go:13
main.main()
	C:/Users/Rongqin/go/src/go-note/go-chan/main/go-select/main/main.go:14 +0x6e

```
```go_select.TestSelect(ch)```中的```ch```
