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
func (Bytes) Lift(s *Store) Bytes {
	size := uint32(s.Stack.Pop())
	offset := uint32(s.Stack.Pop())
	raw, ok := s.Memory.Read(offset, size)
	if !ok {
		s.Error = ErrMemRead
	}
	return Bytes{Offset: offset, Raw: raw}
}

// Lower implements [Lower] interface.
func (v Bytes) Lower(s *Store) {
	ok := s.Memory.Write(v.Offset, v.Raw)
	if !ok {
		s.Error = ErrMemWrite
	}
	size := len(v.Raw)
	s.Stack.Push(Raw(v.Offset))
	s.Stack.Push(Raw(size))
}

// MemoryLift implements [MemoryLift] interface.
func (Bytes) MemoryLift(s *Store, offset uint32) (Bytes, uint32) {
	sp, ok := s.Memory.Read(offset, 8)
	if !ok {
		s.Error = ErrMemRead
		return Bytes{}, 0
	}
	ptr := binary.LittleEndian.Uint32(sp[0:])
	sz := binary.LittleEndian.Uint32(sp[4:])

	raw, ok := s.Memory.Read(ptr, sz)
	if !ok {
		s.Error = ErrMemRead
		return Bytes{}, 0
	}
	return Bytes{Offset: offset, Raw: raw}, sz
}

// MemoryLower implements [MemoryLower] interface.
func (v Bytes) MemoryLower(s *Store, offset uint32) (length uint32) {
	ptrdata := make([]byte, 8)
	binary.LittleEndian.PutUint32(ptrdata[0:], offset+8)
	binary.LittleEndian.PutUint32(ptrdata[4:], uint32(len(v.Raw)))

	ok := s.Memory.Write(offset, ptrdata)
	if !ok {
		s.Error = ErrMemWrite
	}
	ok = s.Memory.Write(offset+8, v.Raw)
	if !ok {
		s.Error = ErrMemWrite
	}
	return uint32(len(v.Raw))
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
func (String) Lift(s *Store) String {
	size := uint32(s.Stack.Pop())
	offset := uint32(s.Stack.Pop())
	raw, ok := s.Memory.Read(offset, size)
	if !ok {
		s.Error = ErrMemRead
	}
	return String{Offset: offset, Raw: string(raw)}
}

// Lower implements [Lower] interface.
func (v String) Lower(s *Store) {
	ok := s.Memory.Write(v.Offset, []byte(v.Raw))
	if !ok {
		s.Error = ErrMemWrite
	}
	size := len(v.Raw)
	s.Stack.Push(Raw(v.Offset))
	s.Stack.Push(Raw(size))
}

// MemoryLift implements [MemoryLift] interface.
func (String) MemoryLift(s *Store, offset uint32) (String, uint32) {
	sp, ok := s.Memory.Read(offset, 8)
	if !ok {
		s.Error = ErrMemRead
		return String{}, 0
	}
	ptr := binary.LittleEndian.Uint32(sp[0:])
	sz := binary.LittleEndian.Uint32(sp[4:])

	raw, ok := s.Memory.Read(ptr, sz)
	if !ok {
		s.Error = ErrMemRead
		return String{}, 0
	}
	return String{Offset: offset, Raw: string(raw)}, sz
}

// MemoryLower implements [MemoryLower] interface.
func (v String) MemoryLower(s *Store, offset uint32) (length uint32) {
	ptrdata := make([]byte, 8)
	binary.LittleEndian.PutUint32(ptrdata[0:], offset+8)
	binary.LittleEndian.PutUint32(ptrdata[4:], uint32(len(v.Raw)))

	ok := s.Memory.Write(offset, ptrdata)
	if !ok {
		s.Error = ErrMemWrite
	}
	ok = s.Memory.Write(offset+8, []byte(v.Raw))
	if !ok {
		s.Error = ErrMemWrite
	}
	return uint32(len(v.Raw))
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
func (ReturnedList[T]) Lift(s *Store) ReturnedList[T] {
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
func (v ReturnedList[T]) Lower(s *Store) {
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

// List wraps a Go slice of any type that implements the [MemoryLiftLower] interface.
// This is the implementation required for the host side of component model functions that pass [cm.List] parameters.
type List[T MemoryLiftLower[T]] struct {
	Offset uint32
	Raw    []T
}

// Unwrap returns the wrapped value.
func (v List[T]) Unwrap() []T {
	return v.Raw
}

// ValueTypes implements [Value] interface.
func (v List[T]) ValueTypes() []ValueType {
	return []ValueType{ValueTypeI32, ValueTypeI32}
}

// Lift implements [Lift] interface.
func (List[T]) Lift(s *Store) List[T] {
	size := uint32(s.Stack.Pop())
	offset := uint32(s.Stack.Pop())
	// empty list
	if size == 0 {
		return List[T]{Offset: offset}
	}
	data := make([]T, size)
	ptr := offset
	var length uint32
	for i := uint32(0); i < size; i++ {
		data[i], length = T.MemoryLift(data[0], s, ptr)
		ptr += length
	}
	return List[T]{Offset: offset, Raw: data}
}

// Lower implements [Lower] interface.
// See https://github.com/WebAssembly/component-model/blob/main/design/mvp/CanonicalABI.md#flattening
// In theory we should re-allocate enough linear memory into which to write the actual data.
func (v List[T]) Lower(s *Store) {
	size := len(v.Raw)
	ptr := v.Offset
	for i := uint32(0); i < uint32(size); i++ {
		length := v.Raw[i].MemoryLower(s, ptr)
		ptr += length
	}
	s.Stack.Push(Raw(v.Offset))
	s.Stack.Push(Raw(size))
}

// MemoryLift implements [MemoryLift] interface.
func (List[T]) MemoryLift(s *Store, offset uint32) (List[T], uint32) {
	sp, ok := s.Memory.Read(offset, 4)
	if !ok {
		s.Error = ErrMemRead
		return List[T]{}, 0
	}
	sz := binary.LittleEndian.Uint32(sp[0:])

	ptr := offset + 4
	data := make([]T, sz)
	var v T
	var length uint32
	for i := uint32(0); i < uint32(sz); i++ {
		data[i], length = v.MemoryLift(s, ptr)
		ptr += length
	}

	return List[T]{Offset: offset, Raw: data}, sz
}

// MemoryLower implements [MemoryLower] interface.
func (v List[T]) MemoryLower(s *Store, offset uint32) (length uint32) {
	sz := len(v.Raw)

	ptr := offset + 4
	for i := uint32(0); i < uint32(sz); i++ {
		length := v.Raw[i].MemoryLower(s, ptr)
		ptr += length
	}

	ptrdata := make([]byte, 4)
	binary.LittleEndian.PutUint32(ptrdata[0:], uint32(sz))

	ok := s.Memory.Write(offset, ptrdata)
	if !ok {
		s.Error = ErrMemWrite
	}

	return uint32(sz)
}

// ListStrings wraps a Go slice of strings.
// This is the implementation required for the host side of component model functions that pass [cm.List] of strings
// as parameters.
// See https://github.com/bytecodealliance/wasm-tools-go/blob/main/cm/list.go
type ListStrings struct {
	Offset uint32
	Raw    []string
}

// Unwrap returns the wrapped value.
func (v ListStrings) Unwrap() []string {
	return v.Raw
}

// ValueTypes implements [Value] interface.
func (v ListStrings) ValueTypes() []ValueType {
	return []ValueType{ValueTypeI32, ValueTypeI32}
}

// Lift implements [Lift] interface.
func (ListStrings) Lift(s *Store) ListStrings {
	size := uint32(s.Stack.Pop())
	offset := uint32(s.Stack.Pop())

	// empty list
	if size == 0 {
		return ListStrings{Offset: offset}
	}

	data := make([]string, size)

	for i := uint32(0); i < size; i++ {
		buf, ok := s.Memory.Read(offset+i*8, 8)
		if !ok {
			s.Error = ErrMemRead
			return ListStrings{Offset: offset, Raw: data}
		}

		ptr := binary.LittleEndian.Uint32(buf[0:])
		sz := binary.LittleEndian.Uint32(buf[4:])

		raw, ok := s.Memory.Read(ptr, sz)
		if !ok {
			s.Error = ErrMemRead
			return ListStrings{Offset: offset, Raw: data}
		}

		data[i] = string(raw)
	}

	return ListStrings{Offset: offset, Raw: data}
}

// Lower implements [Lower] interface.
// See https://github.com/WebAssembly/component-model/blob/main/design/mvp/CanonicalABI.md#flattening
// In theory we should re-allocate enough linear memory into which to write the actual data.
func (v ListStrings) Lower(s *Store) {
	size := uint32(len(v.Raw))
	plen := size * 8

	// write pointers
	for i := uint32(0); i < size; i++ {
		ptrdata := make([]byte, 8)
		binary.LittleEndian.PutUint32(ptrdata[0:], v.Offset+i*8+plen)
		binary.LittleEndian.PutUint32(ptrdata[4:], uint32(len(v.Raw[i])))

		ok := s.Memory.Write(v.Offset+i*8, ptrdata)
		if !ok {
			s.Error = ErrMemRead
			return
		}
	}

	// write the actual strings
	for i, str := range v.Raw {
		ptr := v.Offset + plen + uint32(i)*8

		ok := s.Memory.Write(ptr, []byte(str))
		if !ok {
			s.Error = ErrMemRead
			return
		}
	}

	s.Stack.Push(Raw(v.Offset))
	s.Stack.Push(Raw(size))
}

// Result is the implementation required for the host side of component model functions that return a *[cm.Result] type.
// See https://github.com/bytecodealliance/wasm-tools-go/blob/main/cm/result.go
type Result[Shape MemoryLiftLower[Shape], OK MemoryLiftLower[OK], Err MemoryLiftLower[Err]] struct {
	Offset  uint32
	DataPtr uint32
	OK      OK
	Error   Err
	IsError bool
}

// Unwrap returns the wrapped value.
func (v Result[Shape, OK, Err]) Unwrap() any {
	if v.IsError {
		return v.Error
	}

	return v.OK
}

// ValueTypes implements [Value] interface.
func (v Result[Shape, OK, Err]) ValueTypes() []ValueType {
	return []ValueType{ValueTypeI32}
}

// Lift implements [Lift] interface. Lifting a result is not supported.
func (Result[Shape, OK, Err]) Lift(s *Store) Result[Shape, OK, Err] {
	offset := uint32(s.Stack.Pop())

	var B UInt32
	isError, sz := B.MemoryLift(s, offset)

	if isError > 0 {
		var E Err
		err, _ := E.MemoryLift(s, offset+sz)
		return Result[Shape, OK, Err]{
			IsError: true,
			Error:   err,
			Offset:  offset,
		}
	}

	var T OK
	val, _ := T.MemoryLift(s, offset+sz)
	return Result[Shape, OK, Err]{
		IsError: false,
		OK:      val,
		Offset:  offset,
	}
}

// Lower implements [Lower] interface.
// See https://github.com/WebAssembly/component-model/blob/main/design/mvp/CanonicalABI.md#flattening
// To use this need to have pre-allocated linear memory into which to write the actual data.
func (v Result[Shape, OK, Err]) Lower(s *Store) {
	if v.DataPtr == 0 {
		s.Error = ErrMemWrite
		return
	}

	var isError UInt32
	if v.IsError {
		isError = 1
	}
	sz := isError.MemoryLower(s, v.Offset)

	switch v.IsError {
	case true:
		v.Error.MemoryLower(s, v.Offset+sz)
	case false:
		v.OK.MemoryLower(s, v.Offset+sz)
	}
}

// TODO: fixed-width array
// TODO: CString
