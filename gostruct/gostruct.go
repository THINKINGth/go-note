package gostruct

/*
 * 在调用结构时，结构名同样要大写。
 * Type 是自定义类型的关键字
 */
type Student struct {
	Name string
	Age int
}
func New(name string, age int) * Student{
	return &Student{Name: name, Age: age}
}

func (s Student) Prints_name() string{
	return s.Name
}


