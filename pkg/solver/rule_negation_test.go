package solver

import "testing"

func TestNegation(t *testing.T) {
	c := NewBoardFromString(`
	1oxoo
	ox3xo
	4oo2x
	oxxx1
	ooo0o
	`)

	r := parseTestCases(`
	1____
	__1__
	0__1_
	____0
	___1_
	`)

	n := &Negation{}

	testBoardCheck(t, n, c, r)
}
