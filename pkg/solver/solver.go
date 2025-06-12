package solver

import "fmt"

type Solver struct {
	rules []Rule
	board Board
}

type Result struct{}

func (s *Solver) travelAllCases(board Board, leftUnknown int, hints []Vec2, rule Rule, cb func(board Board)) {
	for _, pos := range hints {
		x, y := pos[0], pos[1]
		cell := board[x][y]
		if cell.Type != CellType_Unknown {
			continue
		}
		originalCell := board[x][y]
		board[x][y] = cc.Empty()
		if leftUnknown > 1 {
			if rule.CanPreCheck() {
				if rule.PreCheck(board, pos) {
					s.travelAllCases(board, leftUnknown-1, hints, rule, cb)
				}
			}
		} else {
			if rule.Check(board, pos) {
				cb(board)
			}
		}

		board[x][y] = cc.Mine()
		if leftUnknown > 1 {
			if rule.CanPreCheck() {
				if rule.PreCheck(board, pos) {
					s.travelAllCases(board, leftUnknown-1, hints, rule, cb)
				}
			}
		} else {
			if rule.Check(board, pos) {
				cb(board)
			}
		}
		board[x][y] = originalCell
	}
}

func (s *Solver) solveSingleRule() (*Result, error) {
	// rule := s.rules[0]
	// hints := rule.GetHintPoints(s.board)

	// if hints != nil {
	// 	if len(hints) > 0 {
	// 		s.travelAllCases(s.board, len(hints), hints, rule, func(board Board) {
	// 			allCases = append(allCases, board.Clone())
	// 		})

	// 	}
	// }

	return nil, fmt.Errorf("not implemented")
}

func (s *Solver) Solve() (*Result, error) {
	if len(s.board) == 0 {
		return nil, fmt.Errorf("board is empty")
	}

	switch len(s.rules) {
	case 0:
		return nil, fmt.Errorf("no rules provided")
	case 1:
		return s.solveSingleRule()
	}

	return nil, fmt.Errorf("multiple rules not supported yet")
}

func (s *Solver) TryShallow(board Board, pos Vec2, cell *Cell) bool {
	x, y := pos[0], pos[1]
	originalCell := board[x][y]
	board[pos[0]][pos[1]] = cell
	defer func() {
		board[x][y] = originalCell
	}()

	for _, rule := range s.rules {
		if !rule.Check(board, pos) {
			board[x][y] = originalCell
			return false
		}
	}
	return true
}

func (s *Solver) TryDeep(board Board, pos Vec2, cell *Cell) bool {
	x, y := pos[0], pos[1]
	originalCell := board[x][y]
	board[pos[0]][pos[1]] = cc.Mine()
	defer func() {
		board[x][y] = originalCell
	}()

	if !s.fillAllUnknownCell(board) {
		return false
	}

	for _, rule := range s.rules {
		if !rule.Check(board, pos) {
			board[x][y] = originalCell
			return false
		}
	}
	return true
}

func (s *Solver) fillAllUnknownCell(b Board) bool {
	nb := b.Clone()
	for x := range b {
		for y := range b[x] {
			if b[x][y].Type == CellType_Unknown {
				if s.TryShallow(b, Vec2{x, y}, cc.Mine()) {
					nb[x][y] = cc.Mine()
					continue
				}
				if s.TryShallow(b, Vec2{x, y}, cc.Empty()) {
					nb[x][y] = cc.Empty()
					continue
				}
				return false
			}
		}
	}
	return true
}
