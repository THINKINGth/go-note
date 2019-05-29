package gofunc

import (
	"fmt"
	"time"
)

func Add (num1 int, num2 int) int {
	return num1 + num2
}

func Sub (num1 int, num2 int) int {
	return num1 - num2
}

/**
 * 多返回值
 */
func Mul_div(num1, num2 int) (int, int){
	return num1 * num2, num1 / num2
}

// 声明返回值，增强了可读性
func Mul_div_other(num1, num2 int) (mul, div int){
	mul = num1 * num2
	div = num1 / num2
	return mul, div
}

/**
 * 参数是函数的函数
 */
 // type typeName func(input1 inputType1, input2 inputType2, [,...]) (result1 resultType1 [,...])
func Formatfunc (f func(n, m int) int, n int, m int) int{
	return f(n, m)
}
// 将复杂签名定义为函数类型
type funcno func(n, m int) int
func Formatfuncn1 (f funcno, n, m int) int{
	return f(n, m)
}

// 变参函数
func Funcnums(arr ...int) []int {
	return arr
}

// 闭包
func Testadd1() func(num1 int) int {
	return func(num1 int) int{
		return num1 + 2
	}
}

func Testadd2(num1 int) func(num2 int) int {
	return func(num2 int) int {
		return num1 + num2
	}
}

// 闭包在环境继承方面
func Getaddnum() func() int{
	i := 0
	return func() int{
		i += 1
		return i
	}
}


// 递归函数

// 参数传递机制
func Addc(num1 int) int {
	num1++
	return num1
}

func AddP(num1 *int) int {
	*num1++
	return *num1
}

// defer 探究

func Defer_think() (int, *int) {
	i := 0
	// i++
	defer func(i int) {
		fmt.Printf("defer3: %d\t%v\t%s\n", i, &i, time.Now())
	}(i)

	defer fmt.Printf("defer2: %d\t%v\t%s\n", i, &i, time.Now())

	defer func() {
		i++
		fmt.Printf("defer1: %d\t%v\t%s\n", i, &i, time.Now())
	}()

	i++
	fmt.Printf("print1: %d\t%v\t%s\n", i, &i, time.Now())
	func () {
		fmt.Printf("print2: %d\t%v\t%s\n", i, &i, time.Now())
	}()
	return i, &i
}

// defer 追踪
func Trace(s string) string {
	fmt.Printf("开始执行：%s\n", s)
	return s
}

func Un(s string) string {
	fmt.Printf("执行结束：%s\n", s)
	return s
}

//被追踪的函数
func Defer_test() {
	defer Un(Trace(" Defer_test\n"))
	fmt.Printf("被追踪的程序\n")
}
// error
// panic
func GenErr() {
	defer func () {
		fmt.Printf("程序执行\n")
		panic("第二个错误")
	}()
	panic("第一个错误")
}
// 错误发生了，再去抓
// 由于recover()函数用于终止错误处理流程，所以在一般情况下，recover()仅在defer语句的“函数”中有效
// 以有效截取错误处理流程，recover()只有在defer的“函数”内直接调用才会终止错误，否则，总是返回nil。
// 如果在没有发生异常的goroutine中明确用recover()函数,会导致该goroutine所属的进程打印异常信息后直接退出。
func ThrowsPanoc(f func ()) {
	defer func() {
		if r := recover(); r != nil {
            fmt.Printf("捕获的异常：%v\n", r)
		}
	}()
	f()
}

// 将数组的引用传给函数
func ArrTest (array *[5]int) {
	fmt.Println(array)
	for i := 0; i < len(array); i++ {
		fmt.Printf("%d ", array[i])
	}
	fmt.Printf("\n")
}
// 将切片的引用传给函数

func SliceTest(array *[]int) {
    fmt.Println(array)
    fmt.Println(*array)
	for i := 0; i < len(*array); i++ {
		fmt.Printf("%d ", (*array)[i])
	}
	fmt.Printf("\n")
}


