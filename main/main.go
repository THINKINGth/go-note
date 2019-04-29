package main

import (
	"fmt"
	"gopratice/gostruct"
	"gopratice/gofunc"
)

func main() {
	var num int
	num = 555
	/*
	 * 调用函数实例
	 */
	fmt.Printf("Hello world !!!  %d\n", num)
	/*
	 * 结构实例
	 */
	var stu = gostruct.Student{Name: "Rongqin", Age: 22}
	fmt.Printf("Name: %s  Age: %d\n", stu.Name, stu.Age)


    /*
     * 常量
     */
	const s0 = 0
	const s1 = 1
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

