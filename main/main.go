package main
/*
 * 学习笔记
 */
import (
	"fmt"
	"go-note/gostruct"
	"go-note/gofunc"
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
    var sli = arr[:2]
    fmt.Printf("%v\n", sli[:])
    sli = make([]int, 5, 10)
    fmt.Printf("%v", sli)
	}

