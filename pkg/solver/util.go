package solver

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

var (
	Near8 = []Vec2{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}
	Near4 = []Vec2{
		{-1, 0},
		{0, -1},
		{0, 1},
		{1, 0},
	}
)

func getRelatedCells(b Board, offsets []Vec2, pos Vec2) []*Cell {
	cells := make([]*Cell, 0, len(offsets))
	for _, offset := range offsets {
		x, y := pos[0]+offset[0], pos[1]+offset[1]
		if b.IsInBounds(x, y) {
			cells = append(cells, b[x][y])
		}
	}
	return cells
}

type BoolMap [][]bool

func NewBoolMap(width, height int) BoolMap {
	m := make(BoolMap, width)
	for x := range m {
		m[x] = make([]bool, height)
	}
	return m
}
