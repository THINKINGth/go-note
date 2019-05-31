package Student

import "go-note/go-struct/Person"

type Student struct {
	Person.Person
	IsNormal bool
	Class string
	Profession string
}

func (student Student) Introduce() string{
	return student.GetName()+ " " +student.Class + " " +student.Profession + "\n"
}