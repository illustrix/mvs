package solver

import "testing"

func TestOutsideCheck(t *testing.T) {
	c1 := NewBoardFromString(`
	ooxoo
	oxoxo
	oooxx
	oxxxo
	ooooo
	`)

	o := &Outside{}

	if !o.Check(c1, Vec2{2, 0}) {
		t.Error("checkMine failed")
	}
}
