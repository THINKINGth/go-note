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
	defer Un(Trace(" Defer_test"))
	fmt.Printf("被追踪的程序\n")
}

