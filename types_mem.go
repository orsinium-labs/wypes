package wypes

type Bytes []byte

func (v Bytes) Unwrap() []byte {
	return []byte(v)
}

// ValueTypes implements [Value] interface.
func (v Bytes) ValueTypes() []ValueType {
	return []ValueType{ValueTypeI32, ValueTypeI32}
}

// Lift implements [Lift] interface.
func (Bytes) Lift(s Store) Bytes {
	offset := s.Stack.Pop()
	size := s.Stack.Pop()
	b, _ := s.Memory.Read(uint32(offset), uint32(size))
	return Bytes(b)
}

type String string

func (v String) Unwrap() string {
	return string(v)
}

// ValueTypes implements [Value] interface.
func (v String) ValueTypes() []ValueType {
	return []ValueType{ValueTypeI32, ValueTypeI32}
}

// Lift implements [Lift] interface.
func (String) Lift(s Store) String {
	offset := s.Stack.Pop()
	size := s.Stack.Pop()
	b, _ := s.Memory.Read(uint32(offset), uint32(size))
	return String(b)
}
