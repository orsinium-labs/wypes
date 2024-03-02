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

func TestInt8(t *testing.T) {
	c := is.NewRelaxed(t)
	stack := newStack()
	store := wypes.Store{Stack: stack}

	// push the value to be checked on the stack
	stack.Push(123)
	is.Equal(c, stack.Size(), 1)

	// lift, stack should be empty
	var i wypes.Int8
	i = i.Lift(store)
	is.Equal(c, stack.Size(), 0)

	// lower, it should put the value on the stack
	i.Lower(store)
	is.Equal(c, stack.Size(), 1)

	// pop from the stack and check the value
	is.Equal(c, stack.Pop(), 123)
	is.Equal(c, stack.Size(), 0)
}
