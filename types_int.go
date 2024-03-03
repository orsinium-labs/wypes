package wypes

// Int8 wraps [int8], a signed 8-bit integer.
type Int8 int8

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

// Int16 wraps [int16], a signed 16-bit integer.
type Int16 int16

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

// Int32 wraps [int32], a signed 32-bit integer.
type Int32 int32

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

// Int64 wraps [int64], a signed 64-bit integer.
type Int64 int64

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
