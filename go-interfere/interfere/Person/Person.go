package Person

type Birthday interface {
    GetDay() string
    GetMonth() string
    GetYear() string
    SetDay(str string)
    SetMonth(str string)
    SetYear(str string)
    GetAge() int
}


type Personer interface {
    Birthday
    GetName() string
    GetSex() string
    SetName(name string)
    SetSex(sex string)

}
