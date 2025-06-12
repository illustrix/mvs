package solver

import "fmt"

type Solver struct {
	rules []Rule
	board Board
}

type Result struct {
	Target    CellType
	Confirmed bool
}

type TravelState struct {
	ShouldStop bool
}

func (s *Solver) travelAllCases(board Board, hints []Vec2, rule Rule, cb func(board Board) (fastStop bool)) (fastStop bool) {
	leftUnknown := len(hints)
	if leftUnknown == 0 {
		return
	}
	pos := hints[0]
	x, y := pos[0], pos[1]
	cell := board[x][y]
	leftHints := hints[1:]
	defer func() {
		board[x][y] = cell
	}()
	board[x][y] = cc.Empty()
	if leftUnknown > 1 {
		if rule.CanPreCheck() {
			if rule.PreCheck(board, pos) {
				if s.travelAllCases(board, leftHints, rule, cb) {
					return true
				}
			}
		}
	} else {
		if rule.Check(board, pos) {
			if cb(board) {
				return true
			}
		}
	}

	board[x][y] = cc.Mine()
	if leftUnknown > 1 {
		if rule.CanPreCheck() {
			if rule.PreCheck(board, pos) {
				if s.travelAllCases(board, leftHints, rule, cb) {
					return true
				}
			}
		}
	} else {
		if rule.Check(board, pos) {
			if cb(board) {
				return true
			}
		}
	}
	return
}

func (s *Solver) test(board Board, pos Vec2, hints []Vec2) *Result {
	x, y := pos[0], pos[1]
	cell := board[x][y]

	canBeEmpty := false
	board[x][y] = cc.Empty()
	s.travelAllCases(board, hints, s.rules[0], func(b Board) (fastStop bool) {
		canBeEmpty = true
		fmt.Printf("Cell %v can be empty, board:\n%s\n", pos, b.String())
		return true
	})

	canBeMine := false
	board[x][y] = cc.Mine()
	s.travelAllCases(board, hints, s.rules[0], func(b Board) (fastStop bool) {
		canBeMine = true
		return true
	})
	board[x][y] = cell
	if canBeEmpty && canBeMine {
		fmt.Printf("Warning: cell %v can be both empty and mine, this is unexpected behavior.\n", pos)
		return nil
	}
	if canBeEmpty {
		return &Result{
			Target:    CellType_Empty,
			Confirmed: true,
		}
	}
	return &Result{
		Target:    CellType_Mine,
		Confirmed: true,
	}
}

func (s *Solver) testEachCell(board Board, hints []Vec2) (Vec2, *Result) {
	for i, pos := range hints {
		newHints := make([]Vec2, len(hints)-1)
		copy(newHints[:i], hints[:i])
		if i < len(hints)-1 {
			copy(newHints[i:], hints[i+1:])
		}
		fmt.Printf("Testing cell %v, remaining hints: %v\n", pos, newHints)
		result := s.test(board, pos, newHints)
		if result != nil {
			return pos, result
		}
	}
	return Vec2{}, nil
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
