package solver

type Rule interface {
	Check(board Board, pos Vec2) bool
}
