package wypes

import "math"

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
	return []ValueType{ValueTypeI32}
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
	return []ValueType{ValueTypeI32}
}

func (Float64) Lift(s Store) Float64 {
	f := math.Float64frombits(s.Stack.Pop())
	return Float64(f)
}

func (v Float64) Lower(s Store) {
	r := math.Float64bits(float64(v))
	s.Stack.Push(Raw(r))
}
