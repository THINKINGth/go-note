package voice

import (
	"go-note/go-struct/Feature"
	"go-note/go-struct/author"
)

type Voice struct {
	name string
	Feature.Feature
	author.Author
}

func (v Voice)GetName() string{
	return v.name
}

func (v Voice)GetAuthor() author.Author{
	return v.Author
}

func (v *Voice)SetName(name string) {
	v.name = name
}

func (v *Voice)SetAuthor(author author.Author) {
	v.Author = author
}
