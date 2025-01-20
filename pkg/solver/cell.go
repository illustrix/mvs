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

func NewCellCreator() *CellCreator {
	return &CellCreator{}
}

func (c *CellCreator) M() *Cell {
	return &Cell{Type: CellType_Mine}
}

func (c *CellCreator) I(num int) *Cell {
	return &Cell{Type: CellType_Num, Num: []int{num}}
}

func (c *CellCreator) O() *Cell {
	return &Cell{Type: CellType_Empty}
}

func (c *CellCreator) U() *Cell {
	return &Cell{Type: CellType_Unknown}
}
