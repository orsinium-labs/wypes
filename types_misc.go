package wypes

import (
	"context"
	"math"
	"time"
)

// Bool wraps [bool].
type Bool bool

// Unwrap returns the wrapped value.
func (v Bool) Unwrap() bool {
	return bool(v)
}

// ValueTypes implements [Value] interface.
func (Bool) ValueTypes() []ValueType {
	return []ValueType{ValueTypeI32}
}

// Lift implements [Lift] interface.
func (Bool) Lift(s Store) Bool {
	return s.Stack.Pop() != 0
}

// Lower implements [Lower] interface.
func (v Bool) Lower(s Store) {
	res := 0
	if v {
		res = 1
	}
	s.Stack.Push(Raw(res))
}

// Float32 wraps [float32].
type Float32 float32

// Unwrap returns the wrapped value.
func (v Float32) Unwrap() float32 {
	return float32(v)
}

// ValueTypes implements [Value] interface.
func (Float32) ValueTypes() []ValueType {
	return []ValueType{ValueTypeF32}
}

// Lift implements [Lift] interface.
func (Float32) Lift(s Store) Float32 {
	f := math.Float32frombits(uint32(s.Stack.Pop()))
	return Float32(f)
}

// Lower implements [Lower] interface.
func (v Float32) Lower(s Store) {
	r := math.Float32bits(float32(v))
	s.Stack.Push(Raw(r))
}

// Float64 wraps [float64].
type Float64 float64

// Unwrap returns the wrapped value.
func (v Float64) Unwrap() float64 {
	return float64(v)
}

// ValueTypes implements [Value] interface.
func (Float64) ValueTypes() []ValueType {
	return []ValueType{ValueTypeF64}
}

// Lift implements [Lift] interface.
func (Float64) Lift(s Store) Float64 {
	f := math.Float64frombits(s.Stack.Pop())
	return Float64(f)
}

// Lower implements [Lower] interface.
func (v Float64) Lower(s Store) {
	res := math.Float64bits(float64(v))
	s.Stack.Push(Raw(res))
}

// Complex64 wraps [complex64].
type Complex64 complex64

// Unwrap returns the wrapped value.
func (v Complex64) Unwrap() complex64 {
	return complex64(v)
}

// ValueTypes implements [Value] interface.
func (Complex64) ValueTypes() []ValueType {
	return []ValueType{ValueTypeF32, ValueTypeF32}
}

// Lift implements [Lift] interface.
func (Complex64) Lift(s Store) Complex64 {
	c := complex(
		math.Float32frombits(uint32(s.Stack.Pop())),
		math.Float32frombits(uint32(s.Stack.Pop())),
	)
	return Complex64(c)
}

// Lower implements [Lower] interface.
func (v Complex64) Lower(s Store) {
	vReal := math.Float32bits(real(v))
	vImag := math.Float32bits(imag(v))
	s.Stack.Push(Raw(vReal))
	s.Stack.Push(Raw(vImag))
}

// Complex128 wraps [complex128].
type Complex128 complex128

// Unwrap returns the wrapped value.
func (v Complex128) Unwrap() complex128 {
	return complex128(v)
}

// ValueTypes implements [Value] interface.
func (Complex128) ValueTypes() []ValueType {
	return []ValueType{ValueTypeF64, ValueTypeF64}
}

// Lift implements [Lift] interface.
func (Complex128) Lift(s Store) Complex128 {
	c := complex(
		math.Float64frombits(uint64(s.Stack.Pop())),
		math.Float64frombits(uint64(s.Stack.Pop())),
	)
	return Complex128(c)
}

// Lower implements [Lower] interface.
func (v Complex128) Lower(s Store) {
	vReal := math.Float64bits(real(v))
	vImag := math.Float64bits(imag(v))
	s.Stack.Push(Raw(vReal))
	s.Stack.Push(Raw(vImag))
}

// Duration wraps [time.Duration].
type Duration time.Duration

// Unwrap returns the wrapped value.
func (v Duration) Unwrap() time.Duration {
	return time.Duration(v)
}

// ValueTypes implements [Value] interface.
func (Duration) ValueTypes() []ValueType {
	return []ValueType{ValueTypeI64}
}

// Lift implements [Lift] interface.
func (Duration) Lift(s Store) Duration {
	return Duration(s.Stack.Pop())
}

// Lower implements [Lower] interface.
func (v Duration) Lower(s Store) {
	s.Stack.Push(Raw(v))
}

// Time wraps [time.Time].
type Time time.Time

// Unwrap returns the wrapped value.
func (v Time) Unwrap() time.Time {
	return time.Time(v)
}

// ValueTypes implements [Value] interface.
func (Time) ValueTypes() []ValueType {
	return []ValueType{ValueTypeI64}
}

// Lift implements [Lift] interface.
func (Time) Lift(s Store) Time {
	return Time(time.Unix(int64(s.Stack.Pop()), 0))
}

// Lower implements [Lower] interface.
func (v Time) Lower(s Store) {
	s.Stack.Push(Raw(time.Time(v).Unix()))
}

// Context wraps [context.Context].
type Context struct{ ctx context.Context }

// Unwrap returns the wrapped value.
func (v Context) Unwrap() context.Context {
	return v.ctx
}

// ValueTypes implements [Value] interface.
func (Context) ValueTypes() []ValueType {
	return []ValueType{}
}

// Lift implements [Lift] interface.
func (Context) Lift(s Store) Context {
	return Context{ctx: s.Context}
}

// Void is a return type of a function that returns nothing.
type Void struct{}

// ValueTypes implements [Value] interface.
func (Void) ValueTypes() []ValueType {
	return []ValueType{}
}

// Lower implements [Lower] interface.
func (Void) Lower(s Store) {}

// Pair wraps two values of arbitrary types.
//
// You can combine multiple pairs to pass more than 2 values at once.
// All values are passed through the stack, not memory.
type Pair[L LiftLower[L], R LiftLower[R]] struct {
	Left  L
	Right R
}

// ValueTypes implements [Value] interface.
func (v Pair[L, R]) ValueTypes() []ValueType {
	types := make([]ValueType, 0, 2)
	types = append(types, v.Left.ValueTypes()...)
	types = append(types, v.Right.ValueTypes()...)
	return types
}

// Lift implements [Lift] interface.
func (Pair[L, R]) Lift(s Store) Pair[L, R] {
	var left L
	var right R
	return Pair[L, R]{
		Left:  left.Lift(s),
		Right: right.Lift(s),
	}
}

// Lower implements [Lower] interface.
func (v Pair[L, R]) Lower(s Store) {
	v.Left.Lower(s)
	v.Right.Lower(s)
}

// HostRef is a reference to a Go object stored on the host side in [Refs].
//
// References created this way are never collected by GC because there is no way
// to know if the wasm module still needs it. So it is important to explicitly clean
// references by calling [HostRef.Drop].
//
// A common usage pattern is to create a reference in one host-defined function,
// return it into the wasm module, and then clean it up in another host-defined function
// caled from wasm when the guest doesn't need the value anymore.
// In this scenario, the latter function accepts HostRef as an argument and calls its
// [HostRef.Drop] method. After that, the reference is removed from [Refs] in the [Store]
// and will be eventually collected by GC.
type HostRef[T any] struct {
	Raw   T
	index uint32
	refs  Refs
}

// Unwrap returns the wrapped value.
func (v HostRef[T]) Unwrap() T {
	return v.Raw
}

// Drop remove the reference from [Refs] in [Store].
//
// Can be called only on lifted references
// (passed as an argument into a host-defined function).
func (v HostRef[T]) Drop() {
	if v.refs != nil {
		v.refs.Drop(v.index)
	}
}

// ValueTypes implements [Value] interface.
func (HostRef[T]) ValueTypes() []ValueType {
	return []ValueType{ValueTypeExternref}
}

// Lift implements [Lift] interface.
func (HostRef[T]) Lift(s Store) HostRef[T] {
	index := uint32(s.Stack.Pop())
	var def T
	raw, found := s.Refs.Get(index, def)
	if !found {
		s.Errors = append(s.Errors, ErrRefNotFound)
	}
	return HostRef[T]{
		Raw:   raw.(T),
		index: index,
		refs:  s.Refs,
	}
}

// Lower implements [Lower] interface.
func (v HostRef[T]) Lower(s Store) {
	var index uint32
	if v.index == 0 {
		index = s.Refs.Put(v.Raw)
	} else {
		index = v.index
		s.Refs.Set(v.index, v.Raw)
	}
	s.Stack.Push(Raw(index))
}
