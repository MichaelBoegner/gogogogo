package main

import (
	"fmt"
)

type Article struct {
	Title  string
	Author string
}

func (a Article) String() string {
	return fmt.Sprintf("The %q article was written by %v.", a.Title, a.Author)
}

type Book struct {
	Title  string
	Author string
	Pages  int
}

func (b Book) String() string {
	return fmt.Sprintf("This book, %q, by %s, is %v pages of complete bullshit", b.Title, b.Author, b.Pages)
}

type Stringer interface {
	String() string
}

func main() {
	a := Article{
		Title:  "Understanding Interfaces in Go",
		Author: "Sammy Sharkson",
	}
	Print(a)
	b := Book{
		Title:  "How to Understand Your Dog",
		Author: "Con Artist McGee",
		Pages:  134,
	}
	Print(b)
}

func Print(s Stringer) {
	fmt.Println(s.String())
}
