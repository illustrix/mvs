package solver

import (
	"fmt"
	"testing"
)

func TestTravelAllCases(t *testing.T) {
	// b0 := NewBoardFromString(`
	// 	_x______
	// 	x4ox____
	// 	x4ox____
	// 	_x__4___
	// 	x332____
	// 	__214___
	// 	_______2
	// 	________
	// `)
	b0 := NewBoardFromString(`
		_____
		_4o__
		_4o__
		_____
		_____
	`)
	rule := &Cross{}
	s0 := &Solver{
		rules: []Rule{rule},
		board: b0,
	}
	hints := rule.GetHintPoints(b0)
	fmt.Printf("Hints: %v\n", len(hints))
	// var c int
	// pointsCases := make(map[string][]int)
	// for _, pos := range hints {
	// 	pointsCases[pos.String()] = []int{0, 0} // [mines, non-mines]
	// }
	// pos, r := s0.testEachCell(b0, hints)
	pos := Vec2{0, 1}                                              // Example position to test
	hints = []Vec2{{1, 4}, {0, 2}, {3, 2}, {1, 0}, {1, 3}, {3, 1}} // Example hints
	r := s0.test(b0, pos, hints)
	fmt.Printf("Test each cell result: %v, %+v\n", pos, r)
	// s0.travelAllCases(b0, hints, rule, func(board Board) (fastStop bool) {
	// 	c += 1
	// 	if c%1000 == 0 {
	// 		fmt.Printf("Processed %d cases\n%s\n", c, board.String())
	// 	}
	// 	for _, pos := range hints {
	// 		x, y := pos[0], pos[1]
	// 		cell := board[x][y]
	// 		t := pointsCases[pos.String()]
	// 		if cell.Type == CellType_Mine {
	// 			t[0] += 1
	// 		} else {
	// 			t[1] += 1
	// 		}
	// 	}
	// 	return
	// })
	// fmt.Printf("Total cases: %d\n", c)
	// fmt.Printf("Points cases: %+v\n", pointsCases)
}
