package main
/*
 * 学习笔记
 */
import (
	"fmt"
	"go-note/gostruct"
	"go-note/gofunc"
	"time"
)

func main() {
	// 函数
	// 一般调用
	num1 := 1
	num2 := 2
	fmt.Printf("add: %d    sub: %d\n", gofunc.Add(num1, num2), gofunc.Sub(num1, num2), )
	fmt.Printf("%d\n", gofunc.Add(5, gofunc.Sub(num1, num2))) // 6
	// 调用函数实例
	// 多返回值的函数中使用空白标识符(_)
	var num0, _ = gofunc.Mul_div_other(num1, num2)
	fmt.Printf("mul: %d\n", num0) // 20
	// 调用标准包fmt的Printf()函数
	// ！！注意函数Printf的首字母大写，在golang中首字母大写代表该变量（函数，结构 ...）是可引用的。
	// golang利用这个规则实现了封装
	fmt.Printf("\n")
	// 函数作为参数
	// 函数是第一类变量，可作为参数传递。只要被调用的函数的返回值个数，返回值类型和返回值的顺序与调用函数所需求的实参是
	//一致的，就可以把这个被调用的函数当做其他函数的参数。
	fmt.Printf("函数作为参数：%d\n", gofunc.Formatfunc(gofunc.Add, num1, num2))
	fmt.Printf("利用类型定义函数，构造通用的函数接口，拥有相同的参数和返回值的函数可以作为参数被调用\n")
	fmt.Printf("gofunc.Add：%d\n", gofunc.Formatfuncn1(gofunc.Sub, num1, num2))
	fmt.Printf("gofunc.Sub：%d\n", gofunc.Formatfuncn1(gofunc.Add, num1, num2))
	// 函数在go 语言中也是一种变量，前面用通过Type来定义它，它的类型就是所有拥有相同参数与相同返回值的一种函数类型（函数签名相同的）
	// 把函数作为类型最大的好处在于可以把这个类型的函数当作值来传递
	//    在写一些通用接口的时候，把函数当作值和类型会非常有用，拥有相同函数签名的函数可以作为参数被调用，通过这种模式可以实现
	//各种各样的逻辑，使得程序变得非常灵活。
    fmt.Printf("\n")
	// 可变参数
	// 在golang中，函数的最后一个参数如果是...type的形式，那这个函数就是一个变参函数，它可以处理变长的参数，而这个长度可以是
	// 0。注意的是变参函数中，无论变参有多少个，它们的类型全部都一样。
	arrs := gofunc.Funcnums(5, 4, 3, 2, 1)
	for i := 0; i < len(arrs); i++ {
		fmt.Printf("第%d个参数: %d\t", i + 1, arrs[i])
	}
	fmt.Printf("\n")
    //    最后是关于可变参数为多种类型的情况，之前的例子都是把变参类型约束为一种类型。但如果希望传递任意类型，可以
    // 指定类型为interface{}，也就是接口。
    // 用interface{} 传递任意类型的数据是Go语言的习惯用法。interface{}是类型安全的，这和C, C++不太一样。

    // 匿名函数和闭包
    // 匿名函数
	fmt.Printf("\n")
    var anony = func(num1, num2 int) int{return num1 + num2}
    fmt.Printf("%d\n", anony(5, 4))
    fmt.Printf("匿名函数: %d", func(num []int) int {
		sum := 0
    	for i := range arrs{
    	    sum += i
		}
		return sum
	}(arrs))
	fmt.Printf("\n")
    // 闭包
    // 匿名函数同样称为闭包（函数式语言的术语），简单来说闭包允许调用定义在其他环境下的变量，
    //   可使得某个函数捕获到一些外部状态，例如函数被创建时的状态，用专业的语言表述就是：一个闭包继承了
    //    函数声明时的作用域。这种状态（作用域的变量）会共享到闭包的环境中，因此这些变量可以在闭包中被操作，
    //    直到被销毁。闭包经常被用作包装函数，预先定义好一个或多个参数以用于包装，另一种常见的应用就是使用闭包
    //    来完成更加简洁的错误检查。
    fmt.Printf("闭包函数 gofunc.Testadd1: %d\n", gofunc.Testadd1()(num1))
	fmt.Printf("闭包函数 gofunc.Testadd2: %d\n", gofunc.Testadd2(num1)(num2))
    // 闭包在环境继承方面
    // 闭包函数保存并积累其中变量的值，不管外部函数是否退出，它都能够继续操作外部函数中的局部变量，这些局部变量可以是参数。
	fmt.Printf("%d\t", gofunc.Getaddnum()())
	fmt.Printf("%d\t", gofunc.Getaddnum()())
	fmt.Printf("%v\t", gofunc.Getaddnum())
	addnum := gofunc.Getaddnum()
	fmt.Printf("%d\t", addnum())
	fmt.Printf("%d\t", addnum())
	fmt.Printf("%d\t", addnum())
	addnum1 := gofunc.Getaddnum()
	fmt.Printf("%d\t", addnum1())
	fmt.Printf("%d\t", addnum1())
	fmt.Printf("%d\t", addnum1())

	fmt.Printf("\n")
	// 函数进阶
	// 参数传递机制
	// golang 的参数传递可以分为按值传递和按引用传递，golang默认按值传递，传递的是参数的副本，函数接收参数副本后，使用变量的
	//过程中可能对副本的值进行更改，但不会影响原来的变量。换句话说，调用函数时修改参数的值，不会影响原来的实参的值，因为数值变化只作用在
	// 副本上。
	a, b := 10, 10
	fmt.Printf("按值传递：\t传递前：%d\t",  a)
	gofunc.Addc(a)
	fmt.Printf("传递后：%d\n", a)
	fmt.Printf("按引用传递：\t传递前：%d\t",  b)
	gofunc.AddP(&b)
	fmt.Printf("传递后：%d\n", b)
	// 传指针最明显的三点好处
	// 第一： 传指针使得多个函数能够操作同一个对象。
	// 第二： 传指针比较轻量级(8B), 毕竟只是传内存地址，可以用指针传递体积大的结构体。
	// 如果传递值，在每次创建的副本上面就会花费相对较多的系统开销（内存和时间）。所以当
	// 要传递大的结构体的时候，用指针是一个明智的选择。一般来说，传递指针（一个32位或者64位的值）的消耗 比传递副本
	// 占用更少的资源。在函数调用的时，像切片（slice）, 字典(map), 接口(interface),通道(channel)这样的引用类型都是
	// 默认使用引用传递的（即使没有显示的指出指针）
	// 不过需要注意的是，golang中的channel，slice, map这三种类型的实现机制类似指针所以可以直接传递，而不用取地址后传递
	// 指针。不过若函数需改变slice的长度，则仍需要取地址传递指针。
	// 第三： 传递指针给函数不但可以节省内存（因为没有复制变量的值），而且赋予了函数直接修改外部变量的能力，所以被修改的变量不需要
	// 使用return 返回
	// ！！！实际开发中传递一个指针很容易引发一些不确定的事情，请万分小心那些可以改变外部变量的函数，迫不得已使用时请添加注释。
	// 以便他人能够清楚这个函数到底做了些什么
	// 尽管如此，传递指针的优势还是十分明显的，当需要在函数内改变一个占用内存比较大的变量时，传递指针可以极大的减少内存，性能优势尤为出色，
	// 所以这把双刃剑要适当使用
	fmt.Printf("\n")
	// defer(延迟)与追踪
	// 可以在函数中添加多个defer语句，当函数执行到最后时（return 语句执行之前），这些defer语句会按照"逆序"执行，最后该函数才退出。
	// 1. defer的用途
	// IO 操作时， 如果遇到错误，需要提前返回，而返回之前应该关闭相应的资源，否则容易造成资源泄露等问题。
	// defer可以统一定义关闭资源的语句，减少重复的代码。
	// 优雅，简洁，减少重复代码
	defer fmt.Printf("\ndefer语句使得这条语句最后才被显示出来3。\n")
	defer fmt.Printf("\ndefer语句使得这条语句最后才被显示出来2。\n")
	defer fmt.Printf("\ndefer语句使得这条语句最后才被显示出来1。\n")
	// 在defer后指定的函数会在函数退出前调用。如果多次调用defer，那么defer采用后进先出模式。
	f, d := gofunc.Defer_think()
	fmt.Printf("return: %d\t%v\t%s\n", f, d, time.Now())
	for i := 0; i < 10; i++ {
		defer fmt.Printf("%d\t", i)
	}
	/* defer 原理：
	 * 在defer中的表达式必须是函数调用，意味着defer后面不能执行操作语句。
	 * 首先 return 的实现逻辑简单来说就是：
	 *  第一步是给返回值赋值（若为有名返回值则直接赋值，若为匿名函数返回值则先声明再赋值）；
	 *  第二步调用RET返回指令并传入返回值，而RET则会检查defer是否存在，若存在就先逆序插播defer语句
	 *  最后RET携带返回值退出函数
	 *  可以看出return 并非是一个原子操作，函数返回值与return返回值并不一定一致。因此，defer，return，返回值三者的顺序应该是：return 最先给
	 * 返回值赋值；接着defer开始执行一些收尾工作；最后RET指令携带返回值退出函数。
	 * 最后需要特别提醒的是，defer声明时会先计算确定参数的值，defer推迟执行的仅是其函数体，因此defer语句位置并非随意，defer的初始化还是受到
	 * 外部影响的。
	 */
	// 跟踪
	gofunc.Defer_test()
	// error
	// 等待学习完接口后学习
	gofunc.ThrowsPanoc(gofunc.GenErr)
	// 复合数据类型
	// 数组
	// 声明方式一：
	ar := [5]int{1, 5, 4, 3, 2}
	// 声明方式二：
	arra := [...] int {5, 4, 3, 2, 1}
	// 声明方式三：
	arrb := [5]int{1:5, 2:4}
	for i := 0; i < len(ar); i++ {
		fmt.Printf("%d %d %d\n", ar[i], arra[i], arrb[i])
	}

	arrp := [5]*int {new(int), new(int)}
	*arrp[0] = 10
	*arrp[1] = 11
	fmt.Printf("%d\n", *arrp[0])
	strt := [5]string {"red", "blue", "green", "yellow", "white"}
	strf := [5]string{}
	strf = strt
	for s := range strt {
		fmt.Printf("%v ", strf[s])
	}
	fmt.Printf("\n")
    // 多维数组
    var arrm = [5][5]int{{5}, 3:{5}}
    for i := 0 ; i < len(arrm); i++{
    	for j := 0; j < len(arrm[0]); j++{
    		fmt.Printf("%d\t", arrm[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
	for i := range arrm{
        for j := range arrm[0]{
        	fmt.Printf("%d\t", arrm[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
	gofunc.Arr_test(&arra)
	/*
	 * 结构实例
	 */
	var stu = gostruct.Student{Name: "Golang", Age: 22}
	fmt.Printf("Name: %s  Age: %d\n", stu.Name, stu.Age)
    /*
     * 常量
     */
	const(
		s2 = iota
		s3
		s4
	)
	fmt.Printf("%d\n", s4)
	/*
	 * 字符串
	 */
	var strs string = "This is golang."
	var byte_strs = []byte(strs)
	for i := 0; i < len(byte_strs); i++ {
		fmt.Printf("%c", byte_strs[i])
	}
	strs = strs[5:]
	fmt.Printf("\n%s\n", strs)
	/*
	 * 指针
	 */
	var student = gostruct.Student{Name: "point", Age: 0}
	var name = &student
	fmt.Printf("%s\n", name.Name)
	var numz int = 15
	var p * int = &numz
	fmt.Printf("%d\n", *p)
    /*
     * 数组
    */
    var arr = []int{5, 3, 2, 2, 7}
    for i := 0; i < len(arr); i++ {
    	fmt.Printf("%d ", arr[i])
	}
	fmt.Printf("\n")
	gofunc.Insorting(arr)
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d ", arr[i])
	}
    /*
     *  切片
     */
	defer fmt.Printf("\ndefer语句使得这条语句最后才被显示出来0。\n")
    var sli = arr[:2]
    fmt.Printf("%v\n", sli[:])
    sli = make([]int, 5, 10)
    fmt.Printf("%v", sli)
	}

