package solver

import "testing"

func TestNewBoardFromString(t *testing.T) {
	c := FormatBoardString(`
	ooxo1
	oxoxo
	o?oxx
	oxxxo
	ooooo
	`)

	b := NewBoardFromString(c)
	a := b.String()

	if a != c {
		t.Errorf("Failed to parse board from string\nExpected\n%s\nActual\n%s", c, a)
	}
}
