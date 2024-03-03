package wypes

// UInt8 wraps uint8, 8-bit unsigned integer.
type UInt8 uint8

// Unwrap returns the wrapped value.
func (v UInt8) Unwrap() uint8 {
	return uint8(v)
}

// ValueTypes implements [Value] interface.
func (UInt8) ValueTypes() []ValueType {
	return []ValueType{ValueTypeI32}
}

// Lift implements [Lift] interface.
func (UInt8) Lift(s Store) UInt8 {
	return UInt8(s.Stack.Pop())
}

// Lower implements [Lower] interface.
func (v UInt8) Lower(s Store) {
	s.Stack.Push(Raw(v))
}

// UInt16 wraps uint16, 16-bit unsigned integer.
type UInt16 uint16

// Unwrap returns the wrapped value.
func (v UInt16) Unwrap() uint16 {
	return uint16(v)
}

// ValueTypes implements [Value] interface.
func (UInt16) ValueTypes() []ValueType {
	return []ValueType{ValueTypeI32}
}

// Lift implements [Lift] interface.
func (UInt16) Lift(s Store) UInt16 {
	return UInt16(s.Stack.Pop())
}

// Lower implements [Lower] interface.
func (v UInt16) Lower(s Store) {
	s.Stack.Push(Raw(v))
}

// UInt32 wraps uint32, 32-bit unsigned integer.
type UInt32 uint32

// Unwrap returns the wrapped value.
func (v UInt32) Unwrap() uint32 {
	return uint32(v)
}

// ValueTypes implements [Value] interface.
func (UInt32) ValueTypes() []ValueType {
	return []ValueType{ValueTypeI32}
}

// Lift implements [Lift] interface.
func (UInt32) Lift(s Store) UInt32 {
	return UInt32(s.Stack.Pop())
}

// Lower implements [Lower] interface.
func (v UInt32) Lower(s Store) {
	s.Stack.Push(Raw(v))
}

// UInt64 wraps uint64, 64-bit unsigned integer.
type UInt64 uint64

// Unwrap returns the wrapped value.
func (v UInt64) Unwrap() uint64 {
	return uint64(v)
}

// ValueTypes implements [Value] interface.
func (UInt64) ValueTypes() []ValueType {
	return []ValueType{ValueTypeI64}
}

// Lift implements [Lift] interface.
func (UInt64) Lift(s Store) UInt64 {
	return UInt64(s.Stack.Pop())
}

// Lower implements [Lower] interface.
func (v UInt64) Lower(s Store) {
	s.Stack.Push(Raw(v))
}

// UInt wraps uint, 32-bit unsigned integer.
type UInt uint

// Unwrap returns the wrapped value.
func (v UInt) Unwrap() uint {
	return uint(v)
}

// ValueTypes implements [Value] interface.
func (UInt) ValueTypes() []ValueType {
	return []ValueType{ValueTypeI64}
}

// Lift implements [Lift] interface.
func (UInt) Lift(s Store) UInt {
	return UInt(s.Stack.Pop())
}

// Lower implements [Lower] interface.
func (v UInt) Lower(s Store) {
	s.Stack.Push(Raw(v))
}

// UIntPtr wraps uintptr, pointer-sized unsigned integer.
type UIntPtr uintptr

// Unwrap returns the wrapped value.
func (v UIntPtr) Unwrap() uintptr {
	return uintptr(v)
}

// ValueTypes implements [Value] interface.
func (UIntPtr) ValueTypes() []ValueType {
	return []ValueType{ValueTypeI64}
}

// Lift implements [Lift] interface.
func (UIntPtr) Lift(s Store) UIntPtr {
	return UIntPtr(s.Stack.Pop())
}

// Lower implements [Lower] interface.
func (v UIntPtr) Lower(s Store) {
	s.Stack.Push(Raw(v))
}

// Rune is an alias for [UInt32].
type Rune = UInt32

// Byte is an alias for [UInt8].
type Byte = UInt8
