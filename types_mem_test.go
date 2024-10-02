package wypes_test

import (
	"testing"

	"github.com/orsinium-labs/tinytest/is"
	"github.com/orsinium-labs/wypes"
)

func TestBytes(t *testing.T) {
	c := is.NewRelaxed(t)
	stack := wypes.NewSliceStack(4)
	store := wypes.Store{Stack: stack, Memory: wypes.NewSliceMemory(1024)}

	data := []byte("Hello, World!")
	wypes.Bytes{Raw: data}.Lower(store)

	result := wypes.Bytes{}.Lift(store)
	is.SliceEqual(c, result.Unwrap(), data)
}

func TestListEmpty(t *testing.T) {
	c := is.NewRelaxed(t)
	stack := wypes.NewSliceStack(4)
	store := wypes.Store{Stack: stack, Memory: wypes.NewSliceMemory(1024)}

	wypes.List[wypes.UInt32]{Offset: 64}.Lower(store)

	store.Stack.Push(64)
	list := wypes.List[wypes.UInt32]{}.Lift(store)

	is.SliceEqual(c, list.Unwrap(), []wypes.UInt32{})
}

func TestListUint32(t *testing.T) {
	c := is.NewRelaxed(t)
	stack := wypes.NewSliceStack(4)
	store := wypes.Store{Stack: stack, Memory: wypes.NewSliceMemory(1024)}

	data := []wypes.UInt32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	wypes.List[wypes.UInt32]{Offset: 64, Raw: data, DataPtr: 128}.Lower(store)

	store.Stack.Push(64)
	list := wypes.List[wypes.UInt32]{}.Lift(store)

	is.SliceEqual(c, list.Unwrap(), data)
}

func TestListUint16(t *testing.T) {
	c := is.NewRelaxed(t)
	stack := wypes.NewSliceStack(4)
	store := wypes.Store{Stack: stack, Memory: wypes.NewSliceMemory(1024)}

	data := []wypes.UInt16{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	wypes.List[wypes.UInt16]{Offset: 96, Raw: data, DataPtr: 128}.Lower(store)

	store.Stack.Push(96)
	list := wypes.List[wypes.UInt16]{}.Lift(store)

	is.SliceEqual(c, list.Unwrap(), data)
}

func TestListInt16(t *testing.T) {
	c := is.NewRelaxed(t)
	stack := wypes.NewSliceStack(4)
	store := wypes.Store{Stack: stack, Memory: wypes.NewSliceMemory(1024)}

	data := []wypes.Int16{1, -2, 3, 4, -5, 6, -7, 8, 9, 10}
	wypes.List[wypes.Int16]{Offset: 64, Raw: data, DataPtr: 128}.Lower(store)

	store.Stack.Push(64)
	list := wypes.List[wypes.Int16]{}.Lift(store)

	is.SliceEqual(c, list.Unwrap(), data)
}

func TestListInt32(t *testing.T) {
	c := is.NewRelaxed(t)
	stack := wypes.NewSliceStack(4)
	store := wypes.Store{Stack: stack, Memory: wypes.NewSliceMemory(1024)}

	data := []wypes.Int32{1, -2, 3, 4, -5, 6, -7, 8, 9, 10}
	wypes.List[wypes.Int32]{Offset: 64, Raw: data, DataPtr: 128}.Lower(store)

	store.Stack.Push(64)
	list := wypes.List[wypes.Int32]{}.Lift(store)

	is.SliceEqual(c, list.Unwrap(), data)
}

func TestListFloat32(t *testing.T) {
	c := is.NewRelaxed(t)
	stack := wypes.NewSliceStack(4)
	store := wypes.Store{Stack: stack, Memory: wypes.NewSliceMemory(1024)}

	data := []wypes.Float32{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9, 10.1}
	wypes.List[wypes.Float32]{Raw: data, DataPtr: 128}.Lower(store)

	store.Stack.Push(0)
	list := wypes.List[wypes.Float32]{}.Lift(store)

	is.SliceEqual(c, list.Unwrap(), data)
}
