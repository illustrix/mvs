package solver

import "strings"

type Col []*Cell
type Board []Col

func NewBoard(width, height int) Board {
	board := make(Board, width)
	for x := range board {
		col := make(Col, height)
		board[x] = col
		for y := range board[x] {
			board[x][y] = &Cell{
				Type: CellType_Unknown,
				Num:  nil,
				Pos:  Vec2{x, y},
			}
		}
	}
	return board
}

func (b Board) IsInBounds(x, y int) bool {
	return x >= 0 && x < len(b) && y >= 0 && y < len((b)[x])
}

func (b Board) IsEdge(x, y int) bool {
	return x == 0 || x == len(b)-1 || y == 0 || y == len((b)[x])-1
}

func trimBoardString(m string) []string {
	rawLines := strings.Split(m, "\n")
	lines := make([]string, len(rawLines))
	i := 0
	for _, line := range rawLines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		lines[i] = line
		i++
	}
	return lines[:i]
}

func FormatBoardString(m string) string {
	return strings.Join(trimBoardString(m), "\n")
}

func NewBoardFromString(m string) Board {
	lines := trimBoardString(m)
	b := NewBoard(len(lines[0]), len(lines))
	cc := NewCellCreator()
	for y, line := range lines {
		for x, c := range line {
			switch c {
			case 'o':
				b[x][y] = cc.O()
			case 'x':
				b[x][y] = cc.M()
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				b[x][y] = cc.I(int(c - '0'))
			default:
				b[x][y] = cc.U()
			}
			b[x][y].Pos = Vec2{x, y}
		}
	}
	return b
}

func (b Board) String() string {
	var sb strings.Builder
	for y := 0; y < len(b[0]); y++ {
		for x := 0; x < len(b); x++ {
			switch b[x][y].Type {
			case CellType_Empty:
				sb.WriteByte('o')
			case CellType_Mine:
				sb.WriteByte('x')
			case CellType_Num:
				sb.WriteByte(byte('0' + b[x][y].Num[0]))
			default:
				sb.WriteByte('?')
			}
		}
		if y < len(b[0])-1 {
			sb.WriteByte('\n')
		}
	}
	return sb.String()
}
