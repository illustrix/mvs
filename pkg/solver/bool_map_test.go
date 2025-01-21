package solver

import "testing"

func TestBoolMap(t *testing.T) {
	m := NewBoolMap(100, 100)
	for x := range m {
		for y := range m[x] {
			if m[x][y] {
				t.Errorf("Expected false at %d, %d", x, y)
			}
		}
	}
}
