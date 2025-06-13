package solver

import "testing"

func TestCrossPreCheckNonMine(t *testing.T) {
	rule := &Cross{}
	var board Board
	board = NewBoardFromString(`
		_____
		o4o__
		_4o__
		_____
		_____
	`)
	if rule.preCheckNonMine(board, Vec2{0, 1}) {
		t.Errorf("Expected preCheckNonMine to return false, but got true")
	}
	board = NewBoardFromString(`
		_____
		o4o__
		x4ox_
		_____
		_o___
	`)
	if rule.preCheckNonMine(board, Vec2{1, 0}) {
		t.Errorf("Expected preCheckNonMine to return false, but got true")
	}
	board = NewBoardFromString(`
		_____
		_4o__
		_4o__
		_____
		_____
	`)
	if !rule.preCheckNonMine(board, Vec2{2, 1}) {
		t.Errorf("Expected preCheckNonMine to return true, but got false")
	}
}
