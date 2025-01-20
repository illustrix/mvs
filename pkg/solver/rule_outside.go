package solver

type Outside struct {
	board     Board
	traveled  BoolMap
	hasWayOut BoolMap
}

var _ Rule = &Outside{}

func (o *Outside) Check(b Board, pos Vec2) bool {
	o.board = b
	return o.checkMine(pos)
}

func (o *Outside) init() {
	o.traveled = NewBoolMap(len(o.board), len(o.board[0]))
	o.hasWayOut = NewBoolMap(len(o.board), len(o.board[0]))
}

func (o *Outside) dfsForMine(pos Vec2, depth int) bool {
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
			o.hasWayOut[pos[0]][pos[1]] = true
			if depth == 0 {
				return true
			}
			continue
		}
		if cell.Type == CellType_Mine {
			if o.traveled[x][y] {
				continue
			}
			if o.board.IsEdge(x, y) {
				o.hasWayOut[x][y] = true
				if depth == 0 {
					return true
				}
			}
			if !o.traveled[x][y] {
				o.traveled[x][y] = true
				o.dfsForMine(cell.Pos, depth+1)
				o.traveled[x][y] = false
			}
		} else {
			o.traveled[x][y] = true
		}
	}

	return false
}

func (o *Outside) checkMine(pos Vec2) bool {
	o.init()
	return o.dfsForMine(pos, 0)
}
