package go_struct

import (
	"fmt"
)

type Recreation struct {
    Name string
    Classify string
    Kinds string
}

func (rec Recreation) Print() {
	fmt.Printf("%s\n%s\n%s\n", rec.Name, rec.Classify, rec.Kinds)
}

type RecreationMovie struct {
	Recreation
}

type Main interface {
	Introduce() string
}