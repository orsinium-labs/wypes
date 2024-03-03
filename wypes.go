package wypes

import (
	"context"
)

type Raw = uint64
type Addr = uint32
type ValueType = byte

const (
	// ValueTypeI32 is a 32-bit integer.
	ValueTypeI32 ValueType = 0x7f
	// ValueTypeI64 is a 64-bit integer.
	ValueTypeI64 ValueType = 0x7e
	// ValueTypeF32 is a 32-bit floating point number.
	ValueTypeF32 ValueType = 0x7d
	// ValueTypeF64 is a 64-bit floating point number.
	ValueTypeF64 ValueType = 0x7c
)

type Memory interface {
	Read(offset Addr, count uint32) ([]byte, bool)
	Write(offset Addr, v []byte) bool
}

type Stack interface {
	Push(Raw)
	Pop() Raw
}

// SliceStack adapts a slice of raw values into a [Stack].
type SliceStack []uint64

func NewSliceStack(cap int) *SliceStack {
	s := make(SliceStack, 0, cap)
	return &s
}

func (s *SliceStack) Push(v uint64) {
	*s = append(*s, v)
}

func (s *SliceStack) Pop() uint64 {
	idx := len(*s) - 1
	v := (*s)[idx]
	*s = (*s)[:idx]
	return v
}

func (s *SliceStack) Len() int {
	return len(*s)
}

type Store struct {
	Memory  Memory
	Stack   Stack
	Refs    map[uint32]any
	Context context.Context
}

type Value interface {
	ValueTypes() []ValueType
}

type Lift[T any] interface {
	Value
	Lift(Store) T
}

type Lower interface {
	Value
	Lower(Store)
}

type LiftLower[T any] interface {
	Lift[T]
	Lower
}

type Modules map[string]Module

type Module map[string]HostFunc
