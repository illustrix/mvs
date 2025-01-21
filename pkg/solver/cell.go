package solver

type CellType byte

const (
	CellType_Num     CellType = 'i'
	CellType_Mine    CellType = 'x'
	CellType_Unknown CellType = '?'
	CellType_Empty   CellType = 'o'
)

type Cell struct {
	Type CellType
	Num  []int
	Pos  Vec2
}

type CellCreator struct{}

func (c *CellCreator) Mine() *Cell {
	return &Cell{Type: CellType_Mine}
}

func (c *CellCreator) Int(num int) *Cell {
	return &Cell{Type: CellType_Num, Num: []int{num}}
}

func (c *CellCreator) Empty() *Cell {
	return &Cell{Type: CellType_Empty}
}

func (c *CellCreator) Unknown() *Cell {
	return &Cell{Type: CellType_Unknown}
}

var cc = &CellCreator{}
