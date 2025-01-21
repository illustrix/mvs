package solver

type Outside struct {
	board     Board
	traveled  BoolMap
	hasWayOut BoolMap
}

var _ Rule = &Outside{}

func (o *Outside) Check(b Board, pos Vec2) bool {
	o.board = b
	o.init()
	x, y := pos[0], pos[1]
	if b[x][y].Type == CellType_Mine {
		return o.checkMine(pos)
	}
	return o.checkNonMine(pos)
}

func (o *Outside) init() {
	o.traveled = NewBoolMap(len(o.board), len(o.board[0]))
	o.hasWayOut = NewBoolMap(len(o.board), len(o.board[0]))
}

func (o *Outside) doesMineHasWayOut(pos Vec2) bool {
	x, y := pos[0], pos[1]

	if o.hasWayOut[x][y] {
		return true
	}

	if o.board.IsEdge(x, y) {
		o.hasWayOut[x][y] = true
		return true
	}

	cells := getRelatedCells(o.board, Near4, pos)

	for _, cell := range cells {
		x, y := cell.Pos[0], cell.Pos[1]
		if o.hasWayOut[x][y] {
			return true
		}
		if cell.Type == CellType_Mine {
			if o.traveled[x][y] {
				continue
			}
			if !o.traveled[x][y] {
				o.traveled[x][y] = true
				hasWayOut := o.doesMineHasWayOut(cell.Pos)
				if hasWayOut {
					return true
				}
				o.traveled[x][y] = false
			}
		} else {
			o.traveled[x][y] = true
		}
	}

	return false
}

func (o *Outside) checkMine(pos Vec2) bool {
	return o.doesMineHasWayOut(pos)
}

func (o *Outside) travelAllNonMine(pos Vec2) {
	x, y := pos[0], pos[1]
	o.traveled[x][y] = true
	cells := getRelatedCells(o.board, Near4, pos)
	for _, cell := range cells {
		if cell.Type == CellType_Mine {
			continue
		}
		x, y := cell.Pos[0], cell.Pos[1]
		if o.traveled[x][y] {
			continue
		}
		o.travelAllNonMine(cell.Pos)
	}
}

func (o *Outside) checkNonMine(pos Vec2) bool {
	o.travelAllNonMine(pos)
	for x, col := range o.board {
		for y, cell := range col {
			if cell.Type != CellType_Mine && !o.traveled[x][y] {
				return false
			}
		}
	}
	return true
}
