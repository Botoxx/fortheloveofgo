package bookstore_test

import (
	"bookstore"
	"testing"
)

func TestBook(t *testing.T) {
	t.Parallel()
	_ = bookstore.Book{
		Title:  "Ibike Blue Book",
		Author: "John Cena",
		Copies: 3,
	}
}
