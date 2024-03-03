package wypes

// Bytes wraps a slice of bytes.
type Bytes []byte

// Unwrap returns the wrapped value.
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

// String wraps [string].
type String string

// Unwrap returns the wrapped value.
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
