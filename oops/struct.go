package main

import "fmt"

type Book struct {
	name string
	author string
	pages int
}

func (b Book) print_details(){
	fmt.Printf("Book %s was written by %s", b.name, b.author)
    fmt.Printf("\nIt contains %d pages.\n", b.pages)
}

func main() {
	book1 := Book{"Harry Potter", "Shivam", 100}
	book1.print_details()

	book1.name = "The Dead"
	book1.pages = 50

	book1.print_details()
}