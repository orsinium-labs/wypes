package wypes

type UInt8 uint8

func (v UInt8) Unwrap() uint8 {
	return uint8(v)
}

func (UInt8) ValueTypes() []ValueType {
	return []ValueType{ValueTypeI32}
}

func (UInt8) Lift(s Store) UInt8 {
	return UInt8(s.Stack.Pop())
}

func (v UInt8) Lower(s Store) {
	s.Stack.Push(Raw(v))
}

type UInt16 uint16

func (v UInt16) Unwrap() uint16 {
	return uint16(v)
}

func (UInt16) ValueTypes() []ValueType {
	return []ValueType{ValueTypeI32}
}

func (UInt16) Lift(s Store) UInt16 {
	return UInt16(s.Stack.Pop())
}

func (v UInt16) Lower(s Store) {
	s.Stack.Push(Raw(v))
}

type UInt32 uint32

func (v UInt32) Unwrap() uint32 {
	return uint32(v)
}

func (UInt32) ValueTypes() []ValueType {
	return []ValueType{ValueTypeI32}
}

func (UInt32) Lift(s Store) UInt32 {
	return UInt32(s.Stack.Pop())
}

func (v UInt32) Lower(s Store) {
	s.Stack.Push(Raw(v))
}

type UInt64 uint64

func (v UInt64) Unwrap() uint64 {
	return uint64(v)
}

func (UInt64) ValueTypes() []ValueType {
	return []ValueType{ValueTypeI64}
}

func (UInt64) Lift(s Store) UInt64 {
	return UInt64(s.Stack.Pop())
}

func (v UInt64) Lower(s Store) {
	s.Stack.Push(Raw(v))
}

type UInt uint

func (v UInt) Unwrap() uint {
	return uint(v)
}

func (UInt) ValueTypes() []ValueType {
	return []ValueType{ValueTypeI64}
}

func (UInt) Lift(s Store) UInt {
	return UInt(s.Stack.Pop())
}

func (v UInt) Lower(s Store) {
	s.Stack.Push(Raw(v))
}

type UIntPtr uintptr

func (v UIntPtr) Unwrap() uintptr {
	return uintptr(v)
}

func (UIntPtr) ValueTypes() []ValueType {
	return []ValueType{ValueTypeI64}
}

func (UIntPtr) Lift(s Store) UIntPtr {
	return UIntPtr(s.Stack.Pop())
}

func (v UIntPtr) Lower(s Store) {
	s.Stack.Push(Raw(v))
}

type Rune = UInt32
type Byte = UInt8
