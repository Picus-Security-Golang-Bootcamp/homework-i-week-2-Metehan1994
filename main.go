package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	bookList := listingBooks("books.json")

	if len(os.Args) == 1 { //condition for running the program without any arguments
		fmt.Println("You did not write any arguments to check the list of books.")
	} else if os.Args[1] == "list" && len(os.Args) == 2 { //condition for list as an only argument
		//Additional argument together with "list" will be not accepted
		list(bookList)
	} else if len(os.Args) > 2 && os.Args[1] == "search" { //condition includes search and additional arguments accounted for book name.
		//Only search argument will be not accepted for searching a book.
		nameOfBook := ""
		for i := 2; i < len(os.Args); i++ {
			nameOfBook += os.Args[i]
			if i != len(os.Args)-1 { //condition to eliminate adding space after final word
				nameOfBook += " "
			}
		}
		search(nameOfBook, bookList)
	} else {
		fmt.Println("You have written an invalid argument.")
	}
}

//listingBooks reads the json file and converts the json formatted data into proper formatted one for Go
func listingBooks(jsonBookList string) map[string]string {
	bookList := map[string]string{}
	contents, err := ioutil.ReadFile(jsonBookList)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(contents, &bookList); err != nil {
		panic(err)
	}
	return bookList
}

//list shows the whole booklist.
func list(bookList map[string]string) {
	for _, book := range bookList {
		fmt.Println(book)
	}
}

//search takes a specific book name and the booklist, and searches whether the book which is asked is available in the booklist or not.
func search(nameOfBook string, bookList map[string]string) {
	bookFound := false
	nameOfBook = strings.ToLower(nameOfBook)
	for _, book := range bookList {
		lowerCaseBook := strings.ToLower(book)
		if nameOfBook == lowerCaseBook {
			bookFound = true
			fmt.Println(book)
		}
	}
	if !bookFound {
		fmt.Println("There is no such a book in the booklist.")
	}
}
