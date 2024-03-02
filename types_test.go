package wypes_test

import (
	"testing"

	"github.com/orsinium-labs/tinytest/is"
	"github.com/orsinium-labs/wypes"
)

type stack struct {
	raw []uint64
}

func newStack() *stack {
	s := make([]uint64, 0, 4)
	return &stack{s}
}

func (s *stack) Push(v uint64) {
	s.raw = append(s.raw, v)
}

func (s *stack) Pop() uint64 {
	idx := len(s.raw) - 1
	v := s.raw[idx]
	s.raw = s.raw[:idx]
	return v
}

func (s *stack) Size() int {
	return len(s.raw)
}

type tripping[T any] interface {
	wypes.Lower[T]
	wypes.Lift[T]
}

func testRoundtrip[T tripping[T]](t *testing.T) {
	c := is.NewRelaxed(t)
	stack := newStack()
	store := wypes.Store{Stack: stack}

	// push the value to be checked on the stack
	stack.Push(123)
	is.Equal(c, stack.Size(), 1)

	// lift, stack should be empty
	var i T
	i = i.Lift(store)
	is.Equal(c, stack.Size(), 0)

	// lower, it should put the value on the stack
	i.Lower(store)
	is.Equal(c, stack.Size(), 1)

	// pop from the stack and check the value
	is.Equal(c, stack.Pop(), 123)
	is.Equal(c, stack.Size(), 0)
}

// Test that lifting and then lowering a value doesn't change the value.
func TestRoundtrip(t *testing.T) {
	t.Run("Int8", testRoundtrip[wypes.Int8])
	t.Run("Int16", testRoundtrip[wypes.Int16])
	t.Run("Int32", testRoundtrip[wypes.Int32])
	t.Run("Int64", testRoundtrip[wypes.Int64])
}

// A static check that all primitive types can be implicitly cast from literals.
func TestAssignLiteral(t *testing.T) {
	var _ wypes.Int8 = 123
	var _ wypes.Int16 = 12377
	var _ wypes.Int32 = 1237777777
	var _ wypes.Int64 = 1237777777777777777
}
