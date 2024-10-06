package wypes

import "encoding/binary"

// Int8 wraps [int8], a signed 8-bit integer.
type Int8 int8

const int8Size = 1

// Unwrap returns the wrapped value.
func (v Int8) Unwrap() int8 {
	return int8(v)
}

// ValueTypes implements [Value] interface.
func (Int8) ValueTypes() []ValueType {
	return []ValueType{ValueTypeI32}
}

// Lift implements [Lift] interface.
func (Int8) Lift(s Store) Int8 {
	return Int8(s.Stack.Pop())
}

// Lower implements [Lower] interface.
func (v Int8) Lower(s Store) {
	s.Stack.Push(Raw(v))
}

// MemoryLift implements [MemoryLift] interface.
func (Int8) MemoryLift(s Store, offset uint32) (Int8, uint32) {
	raw, ok := s.Memory.Read(offset, int8Size)
	if !ok {
		s.Error = ErrMemRead
		return Int8(0), 0
	}

	return Int8(raw[0]), int8Size
}

// MemoryLower implements [MemoryLower] interface.
func (v Int8) MemoryLower(s Store, offset uint32) (length uint32) {
	ok := s.Memory.Write(offset, []byte{byte(v)})
	if !ok {
		s.Error = ErrMemWrite
		return 0
	}

	return int8Size
}

// Int16 wraps [int16], a signed 16-bit integer.
type Int16 int16

const int16Size = 2

// Unwrap returns the wrapped value.
func (v Int16) Unwrap() int16 {
	return int16(v)
}

// ValueTypes implements [Value] interface.
func (Int16) ValueTypes() []ValueType {
	return []ValueType{ValueTypeI32}
}

// Lift implements [Lift] interface.
func (Int16) Lift(s Store) Int16 {
	return Int16(s.Stack.Pop())
}

// Lower implements [Lower] interface.
func (v Int16) Lower(s Store) {
	s.Stack.Push(Raw(v))
}

// MemoryLift implements [MemoryLifter] interface.
func (Int16) MemoryLift(s Store, offset uint32) (Int16, uint32) {
	raw, ok := s.Memory.Read(offset, int16Size)
	if !ok {
		s.Error = ErrMemRead
		return Int16(0), 0
	}

	return Int16(binary.LittleEndian.Uint16(raw)), int16Size
}

// MemoryLower implements [MemoryLower] interface.
func (v Int16) MemoryLower(s Store, offset uint32) (length uint32) {
	data := make([]byte, int16Size)
	binary.LittleEndian.PutUint16(data, uint16(v))
	ok := s.Memory.Write(offset, data)
	if !ok {
		s.Error = ErrMemWrite
		return 0
	}

	return int16Size
}

// Int32 wraps [int32], a signed 32-bit integer.
type Int32 int32

const int32Size = 4

// Unwrap returns the wrapped value.
func (v Int32) Unwrap() int32 {
	return int32(v)
}

// ValueTypes implements [Value] interface.
func (Int32) ValueTypes() []ValueType {
	return []ValueType{ValueTypeI32}
}

// Lift implements [Lift] interface.
func (Int32) Lift(s Store) Int32 {
	return Int32(s.Stack.Pop())
}

// Lower implements [Lower] interface.
func (v Int32) Lower(s Store) {
	s.Stack.Push(Raw(v))
}

// MemoryLift implements [MemoryLifter] interface.
func (Int32) MemoryLift(s Store, offset uint32) (Int32, uint32) {
	raw, ok := s.Memory.Read(offset, int32Size)
	if !ok {
		s.Error = ErrMemRead
		return Int32(0), 0
	}

	return Int32(binary.LittleEndian.Uint32(raw)), int32Size
}

// MemoryLower implements [MemoryLower] interface.
func (v Int32) MemoryLower(s Store, offset uint32) (length uint32) {
	data := make([]byte, int32Size)
	binary.LittleEndian.PutUint32(data, uint32(v))
	ok := s.Memory.Write(offset, data)
	if !ok {
		s.Error = ErrMemWrite
		return 0
	}

	return int32Size
}

// Int64 wraps [int64], a signed 64-bit integer.
type Int64 int64

const int64Size = 8

// Unwrap returns the wrapped value.
func (v Int64) Unwrap() int64 {
	return int64(v)
}

// ValueTypes implements [Value] interface.
func (Int64) ValueTypes() []ValueType {
	return []ValueType{ValueTypeI64}
}

// Lift implements [Lift] interface.
func (Int64) Lift(s Store) Int64 {
	return Int64(s.Stack.Pop())
}

// Lower implements [Lower] interface.
func (v Int64) Lower(s Store) {
	s.Stack.Push(Raw(v))
}

// MemoryLift implements [MemoryLifter] interface.
func (Int64) MemoryLift(s Store, offset uint32) (Int64, uint32) {
	raw, ok := s.Memory.Read(offset, int64Size)
	if !ok {
		s.Error = ErrMemRead
		return Int64(0), 0
	}

	return Int64(binary.LittleEndian.Uint64(raw)), int64Size
}

// MemoryLower implements [MemoryLower] interface.
func (v Int64) MemoryLower(s Store, offset uint32) (length uint32) {
	data := make([]byte, int64Size)
	binary.LittleEndian.PutUint64(data, uint64(v))
	ok := s.Memory.Write(offset, data)
	if !ok {
		s.Error = ErrMemWrite
		return 0
	}

	return int64Size
}

// Int wraps [int], a signed 32-bit integer.
type Int int

// Unwrap returns the wrapped value.
func (v Int) Unwrap() int {
	return int(v)
}

// ValueTypes implements [Value] interface.
func (Int) ValueTypes() []ValueType {
	return []ValueType{ValueTypeI64}
}

// Lift implements [Lift] interface.
func (Int) Lift(s Store) Int {
	return Int(s.Stack.Pop())
}

// Lower implements [Lower] interface.
func (v Int) Lower(s Store) {
	s.Stack.Push(Raw(v))
}

// MemoryLift implements [MemoryLifter] interface.
func (Int) MemoryLift(s Store, offset uint32) (Int, uint32) {
	raw, ok := s.Memory.Read(offset, int64Size)
	if !ok {
		s.Error = ErrMemRead
		return Int(0), 0
	}

	return Int(binary.LittleEndian.Uint64(raw)), int64Size
}

// MemoryLower implements [MemoryLower] interface.
func (v Int) MemoryLower(s Store, offset uint32) (length uint32) {
	data := make([]byte, int64Size)
	binary.LittleEndian.PutUint64(data, uint64(v))
	ok := s.Memory.Write(offset, data)
	if !ok {
		s.Error = ErrMemWrite
		return 0
	}

	return int64Size
}
