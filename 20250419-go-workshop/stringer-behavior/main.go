package main

import "fmt"

type Article struct {
	Title  string
	Author string
}

func (a Article) String() string {
	return fmt.Sprintf("The %q article was written by %s", a.Title, a.Author)
}

type Book struct {
	Title  string
	Author string
	Pages  int
}

func (b Book) String() string {
	return fmt.Sprintf("The %q book was written by %s and has %d pages", b.Title, b.Author, b.Pages)
}

type Stringer interface {
	String() string
}

func main() {
	a := Article{
		Title:  "Understanding Go Interfaces",
		Author: "Sammy the Shark",
	}

	Print(a)

	b := Book{
		Title:  "All About Go",
		Author: "Jenny Dolphin",
		Pages:  25,
	}

	Print(b)
}

func Print(a Stringer) {
	fmt.Println(a.String())
}
