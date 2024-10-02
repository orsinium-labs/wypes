package wypes

import (
	"encoding/binary"
)

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
	raw, ok := s.Memory.Read(offset, size)
	if !ok {
		s.Error = ErrMemRead
	}
	return Bytes{Offset: offset, Raw: raw}
}

// Lower implements [Lower] interface.
func (v Bytes) Lower(s Store) {
	ok := s.Memory.Write(v.Offset, v.Raw)
	if !ok {
		s.Error = ErrMemWrite
	}
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
	raw, ok := s.Memory.Read(offset, size)
	if !ok {
		s.Error = ErrMemRead
	}
	return String{Offset: offset, Raw: string(raw)}
}

// Lower implements [Lower] interface.
func (v String) Lower(s Store) {
	ok := s.Memory.Write(v.Offset, []byte(v.Raw))
	if !ok {
		s.Error = ErrMemWrite
	}
	size := len(v.Raw)
	s.Stack.Push(Raw(v.Offset))
	s.Stack.Push(Raw(size))
}

// ReturnedList wraps a Go slice of any type that supports the [MemoryLiftLower] interface so it can be returned as a List.
// This is the implementation required for the host side of component model functions that return a *[cm.List] type.
// See https://github.com/bytecodealliance/wasm-tools-go/blob/main/cm/list.go
type ReturnedList[T MemoryLiftLower[T]] struct {
	Offset  uint32
	DataPtr uint32
	Raw     []T
}

// Unwrap returns the wrapped value.
func (v ReturnedList[T]) Unwrap() []T {
	return v.Raw
}

// ValueTypes implements [Value] interface.
func (v ReturnedList[T]) ValueTypes() []ValueType {
	return []ValueType{ValueTypeI32}
}

// Lift implements [Lift] interface.
func (ReturnedList[T]) Lift(s Store) ReturnedList[T] {
	offset := uint32(s.Stack.Pop())
	buf, ok := s.Memory.Read(offset, 8)
	if !ok {
		s.Error = ErrMemRead
		return ReturnedList[T]{}
	}

	ptr := binary.LittleEndian.Uint32(buf[0:])
	sz := binary.LittleEndian.Uint32(buf[4:])

	// empty list, probably a return value to be filled in later.
	if ptr == 0 || sz == 0 {
		return ReturnedList[T]{Offset: offset}
	}

	data := make([]T, sz)
	p := ptr
	var length uint32
	for i := uint32(0); i < sz; i++ {
		data[i], length = T.MemoryLift(data[0], s, p)
		p += length
	}

	return ReturnedList[T]{Offset: offset, DataPtr: ptr, Raw: data}
}

// Lower implements [Lower] interface.
// See https://github.com/WebAssembly/component-model/blob/main/design/mvp/CanonicalABI.md#flattening
// To use this need to have pre-allocated linear memory into which to write the actual data.
func (v ReturnedList[T]) Lower(s Store) {
	if v.DataPtr == 0 {
		s.Error = ErrMemWrite
		return
	}

	size := len(v.Raw)

	ptr := v.DataPtr
	for i := uint32(0); i < uint32(size); i++ {
		length := v.Raw[i].MemoryLower(s, ptr)
		ptr += length
	}

	ptrdata := make([]byte, 8)
	binary.LittleEndian.PutUint32(ptrdata[0:], v.DataPtr)
	binary.LittleEndian.PutUint32(ptrdata[4:], uint32(len(v.Raw)))
	s.Memory.Write(v.Offset, ptrdata)
}

// TODO: fixed-width array
// TODO: CString
