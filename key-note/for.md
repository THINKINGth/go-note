## for 循环

> for variable := range variables
>> 首先是一个坑

>> 对于通道

&emsp;&emsp;首先，for variable := range variables不能用于没有关闭的非阻塞通道，不然会造成死锁。
```go

``` 
&emsp;&emsp;关闭后的通道仍然可以读取。

```go
func main() {
	in := make(chan int)
	close(in)
	fmt.Println(<-in)
}
```

>> 关于指针，不多解释

```go
for i := range chanTask{
    fmt.Printf("%p ", &i)
    // 此时，运行do(&i)时， &i指向的是同一个值
    // go do(i)
}
```
```
0xc04204a3e0
0xc04204a3e0
0xc04204a3e0
0xc04204a3e0
0xc04204a3e0
0xc04204a3e0
0xc04204a3e0
0xc04204a3e0
0xc04204a3e0
0xc04204a3e0
```
```
0
```

```go
// 利用chain.Chain实现管道功能
	ch := chain.Chain((chain.Chain(chain.Chain(in))))
	// 启动器
	go func() {
		for i := 0; i < len(in); i++ {
			
			in <- 1
		}
		close(in)
		fmt.Println(<-in)
	}()

    for i := range ch{
    	fmt.Println(i)
	}

```

### 应用

> 管道

>> 利用

```go
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
```

func main() {
	in := make(chan int)
	// 利用chain.Chain实现管道功能
	// 启动器
	go func() {
		if cap(in) == 0{
			in <- 1
		}
		for i := 0; i < cap(in); i++ {
			in <- 1
		}
		close(in)
	}()
	ch := chain.Chain((chain.Chain(chain.Chain(in))))
	for i := range ch{
		fmt.Println(i)
	}
}
```


