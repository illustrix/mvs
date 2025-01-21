package solver

type Solver struct {
	rules []Rule
	board Board
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
