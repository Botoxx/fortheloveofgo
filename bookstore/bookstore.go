package bookstore

import (
	"errors"
)

type Book struct {
	Title  string
	Author string
	Copies int
	Id     int
}

func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("out of stock")
	}
	b.Copies--
	return b, nil

}

func GetAllBooks(catalog []Book) []Book {
	return catalog
}

func GetBook(catalog map[int]Book, id int) Book {

	return catalog[id]
}
