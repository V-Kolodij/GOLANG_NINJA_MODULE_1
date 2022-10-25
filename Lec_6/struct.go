package main

import "fmt"

type Book struct {
	title      string
	author     string
	pageNumber int
	production int
}

// Add method to Book struct

func (b Book) printBookInfo() {
	fmt.Println(b.title, b.author, b.pageNumber, b.production)
}

// 2 types of methods: value receiver & pointer receiver  func (b *Book)
// Add value receiver method to Book struct
func (b Book) getTitle() string {
	return b.title
}

// Add pointer receiver method to Book struct
func (b *Book) setTitle(title string) {
	b.title = title
}

type DumbDatabase struct {
	m map[string]string
}

func NewDumpDatabase() *DumbDatabase {
	return &DumbDatabase{
		m: make(map[string]string),
	}
}

func main() {
	user := struct {
		name   string
		age    int
		sex    string
		weight int
		height int
	}{"Vasya", 23, "Male", 75, 185}
	fmt.Println(user)
	fmt.Printf("%+v\n", user)
	book1 := Book{"Harry Potter", "Arthur Konan Doil", 456, 2002}
	book2 := Book{"Harry Potter and Philosopher's stone", "Arthur Konan Doil", 656, 2005}

	fmt.Printf("%+v\n", book1)
	fmt.Printf("%+v\n", book2)
	fmt.Printf("%+v\n", book2.title)

	book3 := newBook("Book3", "I am", 400, 1999)
	fmt.Printf("%+v\n", book3)

	// as func
	printBookInfo(book1)
	printBookInfo(book2)
	//as method
	book1.printBookInfo()
}

func newBook(title, author string, pageNumber, production int) Book {
	return Book{
		title:      title,
		author:     author,
		pageNumber: pageNumber,
		production: production,
	}
}

// func for print book info (must add in all files were you need this)
func printBookInfo(book Book) {
	fmt.Println(book.title, book.author, book.pageNumber, book.production)
}
