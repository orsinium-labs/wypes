package wypes

type Int8 int8

var (
	_ Lift[Int8]  = Int8(0)
	_ Lower[Int8] = Int8(0)
)

func (Int8) ValueTypes() []ValueType {
	return []ValueType{ValueTypeI32}
}

func (Int8) Lift(s Store) Int8 {
	return Int8(s.Stack.Pop())
}

func (v Int8) Lower(s Store) {
	s.Stack.Push(Raw(v))
}
