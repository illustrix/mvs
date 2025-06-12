package solver

import (
	"fmt"
	"testing"
)

func TestTravelAllCases(t *testing.T) {
	b0 := NewBoardFromString(`
		_x______
		x4ox____
		x4ox____
		_x__4___
		x332____
		__214___
		_______2
		________
	`)
	rule := &Cross{}
	s0 := &Solver{
		rules: []Rule{rule},
		board: b0,
	}
	hints := rule.GetHintPoints(b0)
	fmt.Printf("Hints: %v\n", len(hints))
	var c int
	pointsCases := make(map[string][]int)
	for _, pos := range hints {
		pointsCases[pos.String()] = []int{0, 0} // [mines, non-mines]
	}
	s0.travelAllCases(b0, len(hints), hints, rule, func(board Board) {
		c += 1
		if c%1000 == 0 {
			fmt.Printf("Processed %d cases\n%s\n", c, board.String())
		}
		for _, pos := range hints {
			x, y := pos[0], pos[1]
			cell := board[x][y]
			t := pointsCases[pos.String()]
			if cell.Type == CellType_Mine {
				t[0] += 1
			} else {
				t[1] += 1
			}
		}
	})
	fmt.Printf("Total cases: %d\n", c)
	fmt.Printf("Points cases: %+v\n", pointsCases)
}
