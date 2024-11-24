package bookstore

import (
	"errors"
	"fmt"
)

type Book struct {
	Title           string
	Author          string
	Copies          int
	ID              int
	PriceCents      int
	DiscountPercent int
}

type Catalog map[int]Book

func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("out of stock")
	}
	b.Copies--
	return b, nil

}

func (catalog Catalog) GetAllBooks() []Book {
	result := []Book{}
	for _, b := range catalog {
		result = append(result, b)
	}
	return result
}

func (catalog Catalog) GetBook(ID int) (Book, error) {
	b, ok := catalog[ID]
	if !ok {
		return b, fmt.Errorf("%d is a wrong ID", ID)
	}
	return b, nil
}

func (book Book) NetPriceCents() int {
	netPrice := book.PriceCents * (100 - book.DiscountPercent) / 100
	return netPrice
}
