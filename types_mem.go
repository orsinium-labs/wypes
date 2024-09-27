package wypes

import (
	"bytes"
	"encoding/binary"
	"unsafe"
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

// List wraps a Go slice of any type.
// This is the implementation required for the host side of the component model [cm.List] type.
// See https://github.com/bytecodealliance/wasm-tools-go/blob/main/cm/list.go
type List[T int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | float32 | float64] struct {
	Offset  uint32
	DataPtr uint32
	Raw     []T
}

// Unwrap returns the wrapped value.
func (v List[T]) Unwrap() []T {
	return v.Raw
}

// ValueTypes implements [Value] interface.
func (v List[T]) ValueTypes() []ValueType {
	return []ValueType{ValueTypeI32}
}

// Lift implements [Lift] interface.
func (List[T]) Lift(s Store) List[T] {
	offset := uint32(s.Stack.Pop())
	buf, ok := s.Memory.Read(offset, 8)
	if !ok {
		s.Error = ErrMemRead
		return List[T]{}
	}

	ptr := binary.LittleEndian.Uint32(buf[0:])
	sz := binary.LittleEndian.Uint32(buf[4:])

	// empty list, probably a return value to be filled in later.
	if ptr == 0 || sz == 0 {
		return List[T]{Offset: offset}
	}

	raw, ok := s.Memory.Read(ptr, uint32(sz)*uint32(unsafe.Sizeof(T(0))))
	if !ok {
		s.Error = ErrMemRead
		return List[T]{Offset: offset}
	}

	r := bytes.NewReader(raw)
	data := make([]T, sz)
	binary.Read(r, binary.LittleEndian, data)

	return List[T]{Offset: offset, DataPtr: ptr, Raw: data}
}

// Lower implements [Lower] interface.
// See https://github.com/WebAssembly/component-model/blob/main/design/mvp/CanonicalABI.md#flattening
// To use this need to have pre-allocated linear memory into which to write the actual data.
func (v List[T]) Lower(s Store) {
	if v.DataPtr == 0 {
		s.Error = ErrMemWrite
		return
	}

	data := new(bytes.Buffer)
	binary.Write(data, binary.LittleEndian, v.Raw)
	s.Memory.Write(v.DataPtr, data.Bytes())

	ptrdata := make([]byte, 8)
	binary.LittleEndian.PutUint32(ptrdata[0:], v.DataPtr)
	binary.LittleEndian.PutUint32(ptrdata[4:], uint32(len(v.Raw)))
	s.Memory.Write(v.Offset, ptrdata)
}

// TODO: fixed-width array
// TODO: CString
