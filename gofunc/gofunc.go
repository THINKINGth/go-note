package gofunc

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

