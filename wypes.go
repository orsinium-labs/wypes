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

type Lower[T any] interface {
	Value
	Lower(Store) T
}
