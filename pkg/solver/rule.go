package solver

type CheckFunc func(board Board, pos Vec2) bool

type Rule interface {
	CanPreCheck() bool
	PreCheck(board Board, pos Vec2) bool

	Check(board Board, pos Vec2) bool

	GetHintPoints(board Board) []Vec2
}

type BaseRule struct{}

func (b *BaseRule) CanPreCheck() bool {
	return false
}

func (b *BaseRule) PreCheck(board Board, pos Vec2) bool {
	panic("unimplemented")
}

func (b *BaseRule) Check(board Board, pos Vec2) bool {
	panic("unimplemented")
}

func (b *BaseRule) GetHintPoints(board Board) []Vec2 {
	return nil
}

var _ Rule = &BaseRule{}
