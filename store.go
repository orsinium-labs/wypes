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

// Store provides access for host-defined functions to the runtime data.
type Store struct {
	// Stack is where [Lift] takes the values from and [Lower] puts values to.
	Stack Stack

	// Memory is used by [Lift] and [Lower] of memory-based types,
	// like [Bytes] and [String].
	Memory Memory

	// Refs is used by [HostRef] to pass through the gues module references
	// to complex objects in the host environment that cannot be lowered into wasm.
	Refs Refs

	// Context can be retrieved by the [Context] type.
	Context context.Context
}

// Memory provides access to the linear memory of the wasm runtime.
//
// The interface is compatible with wazero memory.
type Memory interface {
	// Read is used to [Lift] values of memory-backed types, like [Bytes] and [String].
	Read(offset Addr, count uint32) ([]byte, bool)

	// Read is used to [Lower] values of memory-backed types, like [Bytes] and [String].
	Write(offset Addr, v []byte) bool
}

// Refs holds references to Go values that you want to reference from wasm using [HostRef].
type Refs interface {
	Get(idx uint32, def any) (any, bool)
	Set(idx uint32, val any)
	Put(val any) uint32
}

// MapRefs is a simple [Refs] implementation powered by a map.
//
// Must be constructed with [NewMapRefs].
type MapRefs struct {
	raw map[uint32]any
	idx uint32
}

func NewMapRefs() MapRefs {
	return MapRefs{raw: make(map[uint32]any, 0)}
}

func (r MapRefs) Get(idx uint32, def any) (any, bool) {
	val, found := r.raw[idx]
	if !found {
		return def, false
	}
	return val, true
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

// Value is an interface implemented by all the types in wypes.
type Value interface {
	ValueTypes() []ValueType
}

// Lift reads values from [Store] into a native Go value.
type Lift[T any] interface {
	Value
	Lift(Store) T
}

// Lower writes a native Go value into the [Store].
type Lower interface {
	Value
	Lower(Store)
}

// LiftLower is a type that implements both [Lift] and [Lower].
type LiftLower[T any] interface {
	Lift[T]
	Lower
}

// Modules is a collection of host-defined modules.
//
// It maps module names to the module definitions.
type Modules map[string]Module

// Module is a collection of host-defined functions in a module with the same name.
//
// It maps function names to function definitions.
type Module map[string]HostFunc