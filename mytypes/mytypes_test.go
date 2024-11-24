package mytypes_test

import (
	"mytypes"
	"strings"
	"testing"
)

func TestTwice(t *testing.T) {
	t.Parallel()

	i := mytypes.MyInt(5)
	w := mytypes.MyInt(10)
	g := i.Twice()
	if w != g {
		t.Errorf("want %d, got %d", w, g)
	}
}
func TestMyStringLen(t *testing.T) {
	t.Parallel()

	i := mytypes.MyString("kiscica")
	w := 7
	g := i.Len()
	if w != g {
		t.Errorf("want %d, got %d", w, g)
	}
}

func TestStringsBuilder(t *testing.T) {
	t.Parallel()

	var sb strings.Builder

	sb.WriteString("Hello, ")
	sb.WriteString("Gophers!")

	w := "Hello, Gophers!"
	g := sb.String()
	if w != g {
		t.Errorf("want %q, got %q", w, g)
	}
	wLen := 15
	gLen := sb.Len()
	if wLen != gLen {
		t.Errorf("%q: want lenght of %d, got %d", sb.String(), wLen, gLen)
	}
}
