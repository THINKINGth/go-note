package Teacher

import "go-note/go-struct/Person"

type Teacher struct {
	Person.Person
}

func (teacher Teacher) Introduce() string{
	return teacher.GetName() + "\n"
}
