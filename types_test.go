package wypes_test

import (
	"testing"

	"github.com/orsinium-labs/tinytest/is"
	"github.com/orsinium-labs/wypes"
)

func testRoundtrip[T wypes.LiftLower[T]](t *testing.T) {
	c := is.NewRelaxed(t)
	stack := wypes.NewSliceStack(4)
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
	t.Run("Byte", testRoundtrip[wypes.Byte])
	t.Run("Rune", testRoundtrip[wypes.Rune])

	t.Run("Float32", testRoundtrip[wypes.Float32])
	t.Run("Float64", testRoundtrip[wypes.Float64])
	t.Run("Duration", testRoundtrip[wypes.Duration])
	t.Run("Time", testRoundtrip[wypes.Time])
}

func testRoundtripPair[T wypes.LiftLower[T]](t *testing.T) {
	c := is.NewRelaxed(t)
	stack := wypes.NewSliceStack(4)
	store := wypes.Store{Stack: stack}

	// push the values to be checked on the stack
	stack.Push(123)
	stack.Push(79)
	is.Equal(c, stack.Len(), 2)

	// lift, stack should be empty
	var i T
	i = i.Lift(store)
	is.Equal(c, stack.Len(), 0)

	// lower, it should put the values on the stack
	i.Lower(store)
	is.Equal(c, stack.Len(), 2)

	// pop from the stack and check the value
	is.Equal(c, stack.Pop(), 123)
	is.Equal(c, stack.Pop(), 79)
	is.Equal(c, stack.Len(), 0)
}

func TestRoundtripPair(t *testing.T) {
	t.Run("Complex64", testRoundtripPair[wypes.Complex64])
	t.Run("Complex128", testRoundtripPair[wypes.Complex128])
	t.Run("Pair", testRoundtripPair[wypes.Pair[wypes.Int16, wypes.Int32]])
}

// A static check that all primitive types can be implicitly cast from literals.
func TestAssignLiteral(t *testing.T) {
	var _ wypes.Int8 = 123
	var _ wypes.Int16 = 12377
	var _ wypes.Int32 = 1237777777
	var _ wypes.Int64 = 1237777777777777777
	var _ wypes.Int = 1237777777

	var _ wypes.UInt8 = 123
	var _ wypes.UInt16 = 12377
	var _ wypes.UInt32 = 1237777777
	var _ wypes.UInt64 = 1237777777777777777
	var _ wypes.UInt = 1237777777

	var _ wypes.Float32 = 1.5
	var _ wypes.Float64 = 1.5
	var _ wypes.Complex64 = 3.4 + 1.5i
	var _ wypes.Complex128 = 3.4 + 1.5i
	var _ wypes.Bool = true
}

func TestString_Lift(t *testing.T) {
	c := is.NewRelaxed(t)
	stack := wypes.NewSliceStack(4)
	memory := wypes.NewSliceMemory(40)
	ok := memory.Write(3, []byte("hello!"))
	is.True(c, ok)
	store := wypes.Store{Stack: stack, Memory: memory}
	stack.Push(3) // offset
	stack.Push(6) // len
	var typ wypes.String
	val := typ.Lift(store)
	is.Equal(c, val.Unwrap(), "hello!")
}

func TestString_Lower(t *testing.T) {
	c := is.NewRelaxed(t)
	stack := wypes.NewSliceStack(4)
	memory := wypes.NewSliceMemory(40)
	store := wypes.Store{Stack: stack, Memory: memory}

	val1 := wypes.String{
		Offset: 3,
		Raw:    "hello!",
	}
	val1.Lower(store)
	val2 := val1.Lift(store)
	is.Equal(c, val2.Unwrap(), "hello!")
}

func TestBytes_Lift(t *testing.T) {
	c := is.NewRelaxed(t)
	stack := wypes.NewSliceStack(4)
	memory := wypes.NewSliceMemory(40)
	ok := memory.Write(3, []byte("hello!"))
	is.True(c, ok)
	store := wypes.Store{Stack: stack, Memory: memory}
	stack.Push(3) // offset
	stack.Push(6) // len
	var typ wypes.Bytes
	val := typ.Lift(store)
	is.SliceEqual(c, val.Unwrap(), []byte("hello!"))
}

func TestBytes_Lower(t *testing.T) {
	c := is.NewRelaxed(t)
	stack := wypes.NewSliceStack(4)
	memory := wypes.NewSliceMemory(40)
	store := wypes.Store{Stack: stack, Memory: memory}

	val1 := wypes.Bytes{
		Offset: 3,
		Raw:    []byte("hello!"),
	}
	val1.Lower(store)
	val2 := val1.Lift(store)
	is.SliceEqual(c, val2.Unwrap(), []byte("hello!"))
}
