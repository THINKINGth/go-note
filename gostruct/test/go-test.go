package test

import (
	"fmt"
	"reflect"
)

type TestInt struct {
	Num int64
}

func (t *TestInt)Test1(num int64) {
	t.Num = num
}

func (t TestInt)Test2(num int64) {
	t.Num = num
}

func (t TestInt)Test3() {
	fmt.Println("test: ", t.Num, "!!")
}

func (t *TestInt)Test4() {
	fmt.Println("类型：", reflect.TypeOf(t), "\n值：", reflect.ValueOf(t))
	fmt.Println(t.Num)
}

