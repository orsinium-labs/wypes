package wypes

type Bytes []byte

func (v Bytes) ValueTypes() []ValueType {
	return []ValueType{ValueTypeI32, ValueTypeI32}
}

func (Bytes) Lift(s Store) Bytes {
	offset := s.Stack.Pop()
	size := s.Stack.Pop()
	b, _ := s.Memory.Read(uint32(offset), uint32(size))
	return Bytes(b)
}

type String string

func (v String) ValueTypes() []ValueType {
	return []ValueType{ValueTypeI32, ValueTypeI32}
}

func (String) Lift(s Store) String {
	offset := s.Stack.Pop()
	size := s.Stack.Pop()
	b, _ := s.Memory.Read(uint32(offset), uint32(size))
	return String(b)
}