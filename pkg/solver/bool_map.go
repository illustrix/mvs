package solver

import "strings"

type BoolMap [][]bool

func NewBoolMap(width, height int) BoolMap {
	m := make(BoolMap, width)
	for x := range m {
		m[x] = make([]bool, height)
	}
	return m
}

func (b BoolMap) String() string {
	var sb strings.Builder
	for y := range b[0] {
		for x := range b {
			if b[x][y] {
				sb.WriteByte('1')
			} else {
				sb.WriteByte('0')
			}
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}
