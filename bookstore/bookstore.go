package bookstore

import "errors"

type Book struct {
	Title  string
	Author string
	Copies int
}

func Buy(b Book) (Book, error) {
	if b.Copies > 0 {
		b.Copies--
		return b, nil
	}
	return Book{}, errors.New("out of stock")
}
