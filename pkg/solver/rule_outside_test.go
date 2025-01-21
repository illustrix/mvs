package solver

import "testing"

func TestOutsideCheckMine(t *testing.T) {
	c := NewBoardFromString(`
	ooxoo
	oxoxo
	oooxx
	oxxxo
	ooooo
	`)

	r := parseTestCases(`
	__1__
	_0_1_
	___11
	_111_
	_____
	`)

	o := &Outside{}

	testBoardCheck(t, o, c, r)
}

func TestOutsideCheckNonMine(t *testing.T) {
	var c Board
	o := &Outside{}

	c = NewBoardFromString(`
	ooxoo
	oxoxo
	oooxx
	oxxxo
	ooooo
	`)

	if o.Check(c, Vec2{0, 0}) {
		t.Errorf("Check failed for (0, 0) expected: false actual: true")
		t.FailNow()
	}

	c = NewBoardFromString(`
	oxxxo
	ooxoo
	xoxox
	ooxoo
	ooooo
	`)

	if !o.Check(c, Vec2{0, 0}) {
		t.Errorf("Check failed for (0, 0) expected: true actual: false")
		t.Logf("Traveled\n%s", o.traveled)
		t.FailNow()
	}
}
