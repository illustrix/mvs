package solver

type Cross struct {
	BaseRule
}

var _ Rule = &Cross{}

func (r *Cross) CanPreCheck() bool {
	return true
}

func (r *Cross) getScopeMineCount(b Board, pos Vec2) int {
	cells := getRelatedCells(b, Cross8, pos)
	mineCount := 0
	for _, cell := range cells {
		if cell.Type == CellType_Mine {
			mineCount++
		}
	}
	return mineCount
}

func (r *Cross) getScopeUnknownCount(b Board, pos Vec2) int {
	cells := getRelatedCells(b, Cross8, pos)
	nonMineCount := 0
	for _, cell := range cells {
		if cell.Type == CellType_Unknown {
			nonMineCount++
		}
	}
	return nonMineCount
}

func (r *Cross) preCheckMine(b Board, pos Vec2) bool {
	cells := getRelatedCells(b, Cross8, pos)
	for _, cell := range cells {
		if cell.Type != CellType_Num {
			continue
		}
		if !r.preCheckNum(b, cell.Pos) {
			return false
		}
	}
	return true
}

func (r *Cross) preCheckNonMine(b Board, pos Vec2) bool {
	cells := getRelatedCells(b, Cross8, pos)
	for _, cell := range cells {
		if cell.Type != CellType_Num {
			continue
		}
		if r.getScopeUnknownCount(b, cell.Pos) < cell.Num[0]-r.getScopeMineCount(b, cell.Pos) {
			return false
		}
	}
	return true
}

func (r *Cross) preCheckNum(b Board, pos Vec2) bool {
	x, y := pos[0], pos[1]
	c := b[x][y]
	return r.getScopeMineCount(b, pos) <= c.Num[0]
}

func (r *Cross) PreCheck(b Board, pos Vec2) bool {
	x, y := pos[0], pos[1]
	c := b[x][y]
	if c.Type == CellType_Mine {
		return r.preCheckMine(b, pos)
	}
	return r.preCheckNonMine(b, pos)
}

func (r *Cross) checkNumCell(b Board, pos Vec2) bool {
	x, y := pos[0], pos[1]
	c := b[x][y]
	return r.getScopeMineCount(b, pos) == c.Num[0]
}

func (r *Cross) Check(b Board, pos Vec2) bool {
	x, y := pos[0], pos[1]
	c := b[x][y]
	if c.Type == CellType_Num {
		return r.checkNumCell(b, pos)
	}
	if c.Type == CellType_Mine {
		for _, cell := range getRelatedCells(b, Cross8, pos) {
			if cell.Type == CellType_Num && !r.checkNumCell(b, cell.Pos) {
				return false
			}
		}
	}
	return true
}

func (r *Cross) GetHintPoints(b Board) []Vec2 {
	pointsMap := make(map[string]Vec2)
	points := make(map[string]int)
	for x := range b {
		for y := range b[x] {
			if b[x][y].Type != CellType_Num {
				continue
			}
			pos := Vec2{x, y}
			if r.Check(b, pos) {
				continue
			}
			cells := getRelatedCells(b, Cross8, Vec2{x, y})
			for _, cell := range cells {
				if cell.Type == CellType_Unknown {
					key := cell.Pos.String()
					if _, exists := pointsMap[key]; !exists {
						pointsMap[key] = cell.Pos
					}
					if n, exists := points[key]; exists {
						points[key] = n + 1
					} else {
						points[key] = 1
					}
				}
			}
		}
	}
	hintPoints := make([]Vec2, 0)
	for pos, count := range points {
		if count > 1 {
			hintPoints = append(hintPoints, pointsMap[pos])
		}
	}
	// TODO: use union-find to discover connected components
	return hintPoints
}
