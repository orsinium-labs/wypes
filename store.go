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
	// ValueTypeExternref is an externref type.
	ValueTypeExternref ValueType = 0x6f
)

type Store struct {
	Memory  Memory
	Stack   Stack
	Refs    Refs
	Context context.Context
}

type Memory interface {
	Read(offset Addr, count uint32) ([]byte, bool)
	Write(offset Addr, v []byte) bool
}

type Refs interface {
	Get(idx uint32, def any) any
	Set(idx uint32, val any)
	Put(val any) uint32
}

type MapRefs struct {
	raw map[uint32]any
	idx uint32
}

func NewMapRefs() MapRefs {
	return MapRefs{raw: make(map[uint32]any, 0)}
}

func (r MapRefs) Get(idx uint32, def any) any {
	val, found := r.raw[idx]
	if !found {
		return def
	}
	return val
}

func (r MapRefs) Set(idx uint32, val any) {
	r.raw[idx] = val
}

func (r MapRefs) Put(val any) uint32 {
	r.idx += 1

	// skip already used cells
	_, used := r.raw[r.idx]
	for used {
		r.idx += 1
		_, used = r.raw[r.idx]
	}

	r.raw[r.idx] = val
	return r.idx
}

func (r MapRefs) Drop(idx uint32) {
	r.raw[idx] = nil
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
