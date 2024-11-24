package mytypes_test

import (
	"mytypes"
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
