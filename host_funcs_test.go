package wypes_test

import (
	"testing"

	"github.com/orsinium-labs/tinytest/is"
	"github.com/orsinium-labs/wypes"
)

func TestH0(t *testing.T) {
	c := is.NewRelaxed(t)
	stack := wypes.NewSliceStack(4)
	store := wypes.Store{Stack: stack}
	f := wypes.H0(func() wypes.Int { return 13 })
	f.Call(&store)
	is.Equal(c, stack.Pop(), 13)
}

func TestH1(t *testing.T) {
	c := is.NewRelaxed(t)
	stack := wypes.NewSliceStack(4)
	store := wypes.Store{Stack: stack}
	stack.Push(12)
	f := wypes.H1(func(x wypes.Int) wypes.Int {
		return x * 2
	})
	f.Call(&store)
	is.Equal(c, stack.Pop(), 24)
}

func TestH2(t *testing.T) {
	c := is.NewRelaxed(t)
	stack := wypes.NewSliceStack(4)
	store := wypes.Store{Stack: stack}
	stack.Push(18)
	stack.Push(12)
	f := wypes.H2(func(a, b wypes.Int) wypes.Int {
		return a - b
	})
	f.Call(&store)
	is.Equal(c, stack.Pop(), 6)
}
