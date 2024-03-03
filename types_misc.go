package wypes

import (
	"context"
	"math"
	"time"
)

type Bool bool

func (Bool) ValueTypes() []ValueType {
	return []ValueType{ValueTypeI32}
}

func (Bool) Lift(s Store) Bool {
	return s.Stack.Pop() != 0
}

func (v Bool) Lower(s Store) {
	res := 0
	if v {
		res = 1
	}
	s.Stack.Push(Raw(res))
}

type Float32 float32

func (Float32) ValueTypes() []ValueType {
	return []ValueType{ValueTypeF32}
}

func (Float32) Lift(s Store) Float32 {
	f := math.Float32frombits(uint32(s.Stack.Pop()))
	return Float32(f)
}

func (v Float32) Lower(s Store) {
	r := math.Float32bits(float32(v))
	s.Stack.Push(Raw(r))
}

type Float64 float64

func (Float64) ValueTypes() []ValueType {
	return []ValueType{ValueTypeF64}
}

func (Float64) Lift(s Store) Float64 {
	f := math.Float64frombits(s.Stack.Pop())
	return Float64(f)
}

func (v Float64) Lower(s Store) {
	res := math.Float64bits(float64(v))
	s.Stack.Push(Raw(res))
}

type Complex64 complex64

func (Complex64) ValueTypes() []ValueType {
	return []ValueType{ValueTypeF32, ValueTypeF32}
}

func (Complex64) Lift(s Store) Complex64 {
	c := complex(
		math.Float32frombits(uint32(s.Stack.Pop())),
		math.Float32frombits(uint32(s.Stack.Pop())),
	)
	return Complex64(c)
}

func (v Complex64) Lower(s Store) {
	vReal := math.Float32bits(real(v))
	vImag := math.Float32bits(imag(v))
	s.Stack.Push(Raw(vReal))
	s.Stack.Push(Raw(vImag))
}

type Complex128 complex128

func (Complex128) ValueTypes() []ValueType {
	return []ValueType{ValueTypeF64, ValueTypeF64}
}

func (Complex128) Lift(s Store) Complex128 {
	c := complex(
		math.Float64frombits(uint64(s.Stack.Pop())),
		math.Float64frombits(uint64(s.Stack.Pop())),
	)
	return Complex128(c)
}

func (v Complex128) Lower(s Store) {
	vReal := math.Float64bits(real(v))
	vImag := math.Float64bits(imag(v))
	s.Stack.Push(Raw(vReal))
	s.Stack.Push(Raw(vImag))
}

type Duration time.Duration

func (Duration) ValueTypes() []ValueType {
	return []ValueType{ValueTypeI64}
}

func (Duration) Lift(s Store) Duration {
	return Duration(s.Stack.Pop())
}

func (v Duration) Lower(s Store) {
	s.Stack.Push(Raw(v))
}

type Time time.Time

func (Time) ValueTypes() []ValueType {
	return []ValueType{ValueTypeI64}
}

func (Time) Lift(s Store) Time {
	return Time(time.Unix(int64(s.Stack.Pop()), 0))
}

func (v Time) Lower(s Store) {
	s.Stack.Push(Raw(time.Time(v).Unix()))
}

type Context struct{ ctx context.Context }

func (c Context) Unwrap() context.Context {
	return c.ctx
}

func (Context) ValueTypes() []ValueType {
	return []ValueType{ValueTypeI64}
}

func (Context) Lift(s Store) Context {
	return Context{ctx: s.Context}
}

type Pair[L LiftLower[L], R LiftLower[R]] struct {
	Left  L
	Right R
}

func (v Pair[L, R]) ValueTypes() []ValueType {
	types := make([]ValueType, 0, 2)
	types = append(types, v.Left.ValueTypes()...)
	types = append(types, v.Right.ValueTypes()...)
	return types
}

func (Pair[L, R]) Lift(s Store) Pair[L, R] {
	var left L
	var right R
	return Pair[L, R]{
		Left:  left.Lift(s),
		Right: right.Lift(s),
	}
}

func (v Pair[L, R]) Lower(s Store) {
	v.Left.Lower(s)
	v.Right.Lower(s)
}

type HostRef[T any] struct {
	idx uint32
	raw any
}

func (v HostRef[T]) Unwrap() T {
	return v.raw.(T)
}

func (HostRef[T]) ValueTypes() []ValueType {
	return []ValueType{ValueTypeI32}
}

func (HostRef[T]) Lift(s Store) HostRef[T] {
	idx := uint32(s.Stack.Pop())
	var def T
	return HostRef[T]{
		idx: idx,
		raw: s.Refs.Get(idx, def),
	}
}

func (v HostRef[T]) Lower(s Store) {
	s.Refs.Set(v.idx, v.raw)
	s.Stack.Push(Raw(v.idx))
}
