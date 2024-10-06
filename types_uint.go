package wypes

import "encoding/binary"

// UInt8 wraps uint8, 8-bit unsigned integer.
type UInt8 uint8

const uInt8Size = 1

// Unwrap returns the wrapped value.
func (v UInt8) Unwrap() uint8 {
	return uint8(v)
}

// ValueTypes implements [Value] interface.
func (UInt8) ValueTypes() []ValueType {
	return []ValueType{ValueTypeI32}
}

// Lift implements [Lift] interface.
func (UInt8) Lift(s *Store) UInt8 {
	return UInt8(s.Stack.Pop())
}

// Lower implements [Lower] interface.
func (v UInt8) Lower(s *Store) {
	s.Stack.Push(Raw(v))
}

// MemoryLift implements [MemoryLift] interface.
func (UInt8) MemoryLift(s *Store, offset uint32) (UInt8, uint32) {
	raw, ok := s.Memory.Read(offset, uInt8Size)
	if !ok {
		s.Error = ErrMemRead
		return UInt8(0), 0
	}

	return UInt8(raw[0]), uInt8Size
}

// MemoryLower implements [MemoryLower] interface.
func (v UInt8) MemoryLower(s *Store, offset uint32) (length uint32) {
	ok := s.Memory.Write(offset, []byte{byte(v)})
	if !ok {
		s.Error = ErrMemWrite
		return 0
	}

	return uInt8Size
}

// UInt16 wraps uint16, 16-bit unsigned integer.
type UInt16 uint16

const uInt16Size = 2

// Unwrap returns the wrapped value.
func (v UInt16) Unwrap() uint16 {
	return uint16(v)
}

// ValueTypes implements [Value] interface.
func (UInt16) ValueTypes() []ValueType {
	return []ValueType{ValueTypeI32}
}

// Lift implements [Lift] interface.
func (UInt16) Lift(s *Store) UInt16 {
	return UInt16(s.Stack.Pop())
}

// Lower implements [Lower] interface.
func (v UInt16) Lower(s *Store) {
	s.Stack.Push(Raw(v))
}

// MemoryLift implements [MemoryLift] interface.
func (UInt16) MemoryLift(s *Store, offset uint32) (UInt16, uint32) {
	raw, ok := s.Memory.Read(offset, uInt16Size)
	if !ok {
		s.Error = ErrMemRead
		return UInt16(0), 0
	}

	return UInt16(binary.LittleEndian.Uint16(raw)), uInt16Size
}

// MemoryLower implements [MemoryLower] interface.
func (v UInt16) MemoryLower(s *Store, offset uint32) (length uint32) {
	data := make([]byte, uInt16Size)
	binary.LittleEndian.PutUint16(data, uint16(v))
	ok := s.Memory.Write(offset, data)
	if !ok {
		s.Error = ErrMemWrite
		return 0
	}

	return uInt16Size
}

// UInt32 wraps uint32, 32-bit unsigned integer.
type UInt32 uint32

const uInt32Size = 4

// Unwrap returns the wrapped value.
func (v UInt32) Unwrap() uint32 {
	return uint32(v)
}

// ValueTypes implements [Value] interface.
func (UInt32) ValueTypes() []ValueType {
	return []ValueType{ValueTypeI32}
}

// Lift implements [Lift] interface.
func (UInt32) Lift(s *Store) UInt32 {
	return UInt32(s.Stack.Pop())
}

// Lower implements [Lower] interface.
func (v UInt32) Lower(s *Store) {
	s.Stack.Push(Raw(v))
}

// MemoryLift implements [MemoryLift] interface.
func (UInt32) MemoryLift(s *Store, offset uint32) (UInt32, uint32) {
	raw, ok := s.Memory.Read(offset, uInt32Size)
	if !ok {
		s.Error = ErrMemRead
		return UInt32(0), 0
	}

	return UInt32(binary.LittleEndian.Uint32(raw)), uInt32Size
}

// MemoryLower implements [MemoryLower] interface.
func (v UInt32) MemoryLower(s *Store, offset uint32) (length uint32) {
	data := make([]byte, uInt32Size)
	binary.LittleEndian.PutUint32(data, uint32(v))
	ok := s.Memory.Write(offset, data)
	if !ok {
		s.Error = ErrMemWrite
		return 0
	}

	return uInt32Size
}

// UInt64 wraps uint64, 64-bit unsigned integer.
type UInt64 uint64

const uInt64Size = 8

// Unwrap returns the wrapped value.
func (v UInt64) Unwrap() uint64 {
	return uint64(v)
}

// ValueTypes implements [Value] interface.
func (UInt64) ValueTypes() []ValueType {
	return []ValueType{ValueTypeI64}
}

// Lift implements [Lift] interface.
func (UInt64) Lift(s *Store) UInt64 {
	return UInt64(s.Stack.Pop())
}

// Lower implements [Lower] interface.
func (v UInt64) Lower(s *Store) {
	s.Stack.Push(Raw(v))
}

// MemoryLift implements [MemoryLift] interface.
func (UInt64) MemoryLift(s *Store, offset uint32) (UInt64, uint32) {
	raw, ok := s.Memory.Read(offset, uInt64Size)
	if !ok {
		s.Error = ErrMemRead
		return UInt64(0), 0
	}

	return UInt64(binary.LittleEndian.Uint64(raw)), uInt64Size
}

// MemoryLower implements [MemoryLower] interface.
func (v UInt64) MemoryLower(s *Store, offset uint32) (length uint32) {
	data := make([]byte, uInt64Size)
	binary.LittleEndian.PutUint64(data, uint64(v))
	ok := s.Memory.Write(offset, data)
	if !ok {
		s.Error = ErrMemWrite
		return 0
	}

	return uInt64Size
}

// UInt wraps uint, 32-bit unsigned integer.
type UInt uint

const uIntSize = 8

// Unwrap returns the wrapped value.
func (v UInt) Unwrap() uint {
	return uint(v)
}

// ValueTypes implements [Value] interface.
func (UInt) ValueTypes() []ValueType {
	return []ValueType{ValueTypeI64}
}

// Lift implements [Lift] interface.
func (UInt) Lift(s *Store) UInt {
	return UInt(s.Stack.Pop())
}

// Lower implements [Lower] interface.
func (v UInt) Lower(s *Store) {
	s.Stack.Push(Raw(v))
}

// MemoryLift implements [Reader] interface.
func (UInt) MemoryLift(s *Store, offset uint32) (UInt, uint32) {
	raw, ok := s.Memory.Read(offset, uIntSize)
	if !ok {
		s.Error = ErrMemRead
		return UInt(0), 0
	}

	return UInt(binary.LittleEndian.Uint64(raw)), uIntSize
}

// MemoryLower implements [MemoryLower] interface.
func (v UInt) MemoryLower(s *Store, offset uint32) (length uint32) {
	data := make([]byte, uIntSize)
	binary.LittleEndian.PutUint64(data, uint64(v))
	ok := s.Memory.Write(offset, data)
	if !ok {
		s.Error = ErrMemWrite
		return 0
	}

	return uIntSize
}

// UIntPtr wraps uintptr, pointer-sized unsigned integer.
type UIntPtr uintptr

const uIntPtrSize = 8

// Unwrap returns the wrapped value.
func (v UIntPtr) Unwrap() uintptr {
	return uintptr(v)
}

// ValueTypes implements [Value] interface.
func (UIntPtr) ValueTypes() []ValueType {
	return []ValueType{ValueTypeI64}
}

// Lift implements [Lift] interface.
func (UIntPtr) Lift(s *Store) UIntPtr {
	return UIntPtr(s.Stack.Pop())
}

// Lower implements [Lower] interface.
func (v UIntPtr) Lower(s *Store) {
	s.Stack.Push(Raw(v))
}

// MemoryLift implements [MemoryLift] interface.
func (UIntPtr) MemoryLift(s *Store, offset uint32) (UIntPtr, uint32) {
	raw, ok := s.Memory.Read(offset, uIntPtrSize)
	if !ok {
		s.Error = ErrMemRead
		return UIntPtr(0), 0
	}

	return UIntPtr(binary.LittleEndian.Uint64(raw)), uIntPtrSize
}

// MemoryLower implements [MemoryLower] interface.
func (v UIntPtr) MemoryLower(s *Store, offset uint32) (length uint32) {
	data := make([]byte, uIntPtrSize)
	binary.LittleEndian.PutUint64(data, uint64(v))
	ok := s.Memory.Write(offset, data)
	if !ok {
		s.Error = ErrMemWrite
		return 0
	}

	return uIntPtrSize
}

// Rune is an alias for [UInt32].
type Rune = UInt32

// Byte is an alias for [UInt8].
type Byte = UInt8
