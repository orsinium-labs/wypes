package wypes_test

import (
	"testing"

	"github.com/orsinium-labs/tinytest/is"
	"github.com/orsinium-labs/wypes"
)

func newStack() *wypes.SliceStack {
	s := make(wypes.SliceStack, 0, 4)
	return &s
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
	is.Equal(c, stack.Len(), 1)

	// lift, stack should be empty
	var i T
	i = i.Lift(store)
	is.Equal(c, stack.Len(), 0)

	// lower, it should put the value on the stack
	i.Lower(store)
	is.Equal(c, stack.Len(), 1)

	// pop from the stack and check the value
	is.Equal(c, stack.Pop(), 123)
	is.Equal(c, stack.Len(), 0)
}

// Test that lifting and then lowering a value doesn't change the value.
func TestRoundtrip(t *testing.T) {
	t.Run("Int8", testRoundtrip[wypes.Int8])
	t.Run("Int16", testRoundtrip[wypes.Int16])
	t.Run("Int32", testRoundtrip[wypes.Int32])
	t.Run("Int64", testRoundtrip[wypes.Int64])
	t.Run("Int", testRoundtrip[wypes.Int])

	t.Run("UInt8", testRoundtrip[wypes.UInt8])
	t.Run("UInt16", testRoundtrip[wypes.UInt16])
	t.Run("UInt32", testRoundtrip[wypes.UInt32])
	t.Run("UInt64", testRoundtrip[wypes.UInt64])
	t.Run("UInt", testRoundtrip[wypes.UInt])
}

// A static check that all primitive types can be implicitly cast from literals.
func TestAssignLiteral(t *testing.T) {
	var _ wypes.Int8 = 123
	var _ wypes.Int16 = 12377
	var _ wypes.Int32 = 1237777777
	var _ wypes.Int64 = 1237777777777777777
}
