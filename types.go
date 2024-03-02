package wypes

type Int8 int8

var _ Lift[Int8] = Int8(0)

func (Int8) Lift(s Store) Int8 {
	return Int8(s.Stack.Pop())
}
