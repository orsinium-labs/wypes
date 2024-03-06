package wypes

// Bytes wraps a slice of bytes.
//
// The bytes are passed through the linear memory.
// Since the memory is controlled and allocated by the guest module,
// you have to provide the Offset to be able to [Lower] the value into the memory.
// The offset should be obtained from the guest module, either as an explicit
// function argument or by calling its allocator.
type Bytes struct {
	Offset uint32
	Raw    []byte
}

// Unwrap returns the wrapped value.
func (v Bytes) Unwrap() []byte {
	return v.Raw
}

// ValueTypes implements [Value] interface.
func (v Bytes) ValueTypes() []ValueType {
	return []ValueType{ValueTypeI32, ValueTypeI32}
}

// Lift implements [Lift] interface.
func (Bytes) Lift(s Store) Bytes {
	size := uint32(s.Stack.Pop())
	offset := uint32(s.Stack.Pop())
	raw, _ := s.Memory.Read(offset, size)
	return Bytes{Offset: offset, Raw: raw}
}

// Lower implements [Lower] interface.
func (v Bytes) Lower(s Store) {
	_ = s.Memory.Write(v.Offset, v.Raw)
	size := len(v.Raw)
	s.Stack.Push(Raw(v.Offset))
	s.Stack.Push(Raw(size))
}

// String wraps [string].
//
// The string is passed through the linear memory.
// Since the memory is controlled and allocated by the guest module,
// you have to provide the Offset to be able to [Lower] the value into the memory.
// The offset should be obtained from the guest module, either as an explicit
// function argument or by calling its allocator.
type String struct {
	Offset uint32
	Raw    string
}

// Unwrap returns the wrapped value.
func (v String) Unwrap() string {
	return v.Raw
}

// ValueTypes implements [Value] interface.
func (v String) ValueTypes() []ValueType {
	return []ValueType{ValueTypeI32, ValueTypeI32}
}

// Lift implements [Lift] interface.
func (String) Lift(s Store) String {
	size := uint32(s.Stack.Pop())
	offset := uint32(s.Stack.Pop())
	raw, _ := s.Memory.Read(offset, size)
	return String{Offset: offset, Raw: string(raw)}
}

// Lower implements [Lower] interface.
func (v String) Lower(s Store) {
	_ = s.Memory.Write(v.Offset, []byte(v.Raw))
	size := len(v.Raw)
	s.Stack.Push(Raw(v.Offset))
	s.Stack.Push(Raw(size))
}

// TODO: arbitrary slice
// TODO: fixed-width array
// TODO: CString
