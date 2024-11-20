package bookstore_test

import (
	"bookstore"
	"github.com/google/go-cmp/cmp"
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

func TestBuy1(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:  "Spark Joy",
		Author: "Lucy Liu",
		Copies: 2,
	}
	want := 1
	result, err := bookstore.Buy(b)
	if err != nil {
		t.Fatal(err)
	}
	got := result.Copies
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestBuy2(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:  "Spark Joy",
		Author: "Lucy Liu",
		Copies: 0,
	}
	_, err := bookstore.Buy(b)
	if err == nil {
		t.Error("want error buying from zero copies, got nil")
	}
}

func TestGetAllBooks(t *testing.T) {
	t.Parallel()

	catalog := []bookstore.Book{
		{Title: "For the love of Go"},
		{Title: "The power of Go"},
	}
	want := []bookstore.Book{
		{Title: "For the love of Go"},
		{Title: "The power of Go"},
	}
	got := bookstore.GetAllBooks(catalog)

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBook(t *testing.T) {
	t.Parallel()
	catalog := map[int]bookstore.Book{
		1: {Title: "For the love of Go", Id: 1},
		2: {Title: "The power of Go", Id: 2},
	}
	want := bookstore.Book{
		Title: "The power of Go", Id: 2,
	}
	got := bookstore.GetBook(catalog, 2)

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
