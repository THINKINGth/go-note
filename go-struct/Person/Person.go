package Person

import  "go-note/go-struct/Birthday"

type Person struct {
	name string
	sex string
	Birthday.Birthday
}

func New(name string, sex string, day, month, year string) *Person {
	var bir Birthday.Birthday
	bir.SetDay(day)
	bir.SetMonth(month)
	bir.SetYear(year)
	return &Person{name, sex, bir}
}

func (person Person) New(name string, sex string, day, month, year interface{}) *Person{
	return &Person{name, sex, *Birthday.New(day, month, year)}
}

func (person Person) GetName() string{
	return person.name
}

func (person Person) GetSex() string {
	return person.sex
}

func (person *Person) SetName(name string) {
	person.name = name
}

func (person *Person) SetSex(sex string) {
	person.sex = sex
}

