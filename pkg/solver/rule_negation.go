package solver

type Negation struct {
	BaseRule
}

var _ Rule = &Negation{}

func (n *Negation) IsColored(pos Vec2) bool {
	return pos[0]%2 != pos[1]%2
}

func (n *Negation) Check(b Board, pos Vec2) bool {
	x, y := pos[0], pos[1]
	c := b[x][y]
	cells := getRelatedCells(b, Near8, c.Pos)
	colored := 0
	uncolored := 0
	for _, cell := range cells {
		if cell.Type == CellType_Mine {
			if n.IsColored(cell.Pos) {
				colored++
			} else {
				uncolored++
			}
		}
	}
	return abs(colored-uncolored) == c.Num[0]
}
