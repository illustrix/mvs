package solver

import "fmt"

type Vec2 [2]int

var _ fmt.Stringer = Vec2{}

func (v Vec2) String() string {
	return fmt.Sprintf("(%d,%d)", v[0], v[1])
}
