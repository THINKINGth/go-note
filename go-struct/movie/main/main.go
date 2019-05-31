package main

import (
	"go-note/go-struct/movie"
	"go-note/go-struct/Person"
	"go-note/go-struct/Feature"
	"go-note/go-struct/author"
	"go-note/go-interfere/interfere/Prints"
)

func main(){
	mv := movie.New(author.Author{*Person.New("1", "女", "13", "5", "1993")},
	"神圣的战争", Feature.Feature{})
	var pt Prints.Printer
	pt = mv
	pt.PrintInf()
	bir := *Person.Person{}.New("2", "女", 14, 6, 1996)
	mvs := movie.New(author.Author{bir}, "红军最强大", Feature.Feature{})
	mvs.PrintInf()
}

