package main


import (
	ClsPerson "go-note/go-struct/Person"
	InfPersons "go-note/go-interfere/interfere/Person"
	"fmt"
)

func main() {
	var person = ClsPerson.Person{}
	person.SetDay("13")
	person.SetMonth("5")
	person.SetYear("1993")
	person.SetName("高橋しょう子")
	person.SetSex("女")
    var persons [10] InfPersons.Personer
    persons[0] = ClsPerson.New("相泽南","女", "14", "6", "1996")
    persons[1] = ClsPerson.New("三上悠亚", "女", "16", "8", "1993")
    persons[2] = ClsPerson.New("桃乃木香奈",  "女", "24", "12", "1996")
    persons[3] = &person
	fmt.Println(persons[0].GetName())
	fmt.Println(persons[1].GetName())
    fmt.Println(persons[2].GetName())
    fmt.Println(persons[3].GetName())
}
