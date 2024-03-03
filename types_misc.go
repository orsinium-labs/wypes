package wypes

import (
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
	r := 0
	if v {
		r = 1
	}
	s.Stack.Push(Raw(r))
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
	r := math.Float64bits(float64(v))
	s.Stack.Push(Raw(r))
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
	r := math.Float32bits(real(v))
	i := math.Float32bits(imag(v))
	s.Stack.Push(Raw(r))
	s.Stack.Push(Raw(i))
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
	r := math.Float64bits(real(v))
	i := math.Float64bits(imag(v))
	s.Stack.Push(Raw(r))
	s.Stack.Push(Raw(i))
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

type Pair[L LiftLower[L], R LiftLower[R]] struct {
	Left  L
	Right R
}

func (p Pair[L, R]) ValueTypes() []ValueType {
	r := make([]ValueType, 0, 2)
	r = append(r, p.Left.ValueTypes()...)
	r = append(r, p.Right.ValueTypes()...)
	return r
}

func (Pair[L, R]) Lift(s Store) Pair[L, R] {
	var l L
	var r R
	return Pair[L, R]{
		Left:  l.Lift(s),
		Right: r.Lift(s),
	}
}

func (v Pair[L, R]) Lower(s Store) {
	v.Left.Lower(s)
	v.Right.Lower(s)
}
