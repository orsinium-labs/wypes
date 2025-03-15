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
	wypes.Bytes{Raw: data}.Lower(&store)

	result := wypes.Bytes{}.Lift(&store)
	is.SliceEqual(c, result.Unwrap(), data)
}

func TestReturnedListEmpty(t *testing.T) {
	c := is.NewRelaxed(t)
	stack := wypes.NewSliceStack(4)
	store := wypes.Store{Stack: stack, Memory: wypes.NewSliceMemory(1024)}

	wypes.ReturnedList[wypes.UInt32]{Offset: 64}.Lower(&store)

	store.Stack.Push(64)
	list := wypes.ReturnedList[wypes.UInt32]{}.Lift(&store)

	is.SliceEqual(c, list.Unwrap(), []wypes.UInt32{})
}

func TestReturnedListUint32(t *testing.T) {
	c := is.NewRelaxed(t)
	stack := wypes.NewSliceStack(4)
	store := wypes.Store{Stack: stack, Memory: wypes.NewSliceMemory(1024)}

	data := []wypes.UInt32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	wypes.ReturnedList[wypes.UInt32]{Offset: 64, Raw: data, DataPtr: 128}.Lower(&store)

	store.Stack.Push(64)
	list := wypes.ReturnedList[wypes.UInt32]{}.Lift(&store)

	is.SliceEqual(c, list.Unwrap(), data)
}

func TestReturnedListUint16(t *testing.T) {
	c := is.NewRelaxed(t)
	stack := wypes.NewSliceStack(4)
	store := wypes.Store{Stack: stack, Memory: wypes.NewSliceMemory(1024)}

	data := []wypes.UInt16{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	wypes.ReturnedList[wypes.UInt16]{Offset: 96, Raw: data, DataPtr: 128}.Lower(&store)

	store.Stack.Push(96)
	list := wypes.ReturnedList[wypes.UInt16]{}.Lift(&store)

	is.SliceEqual(c, list.Unwrap(), data)
}

func TestReturnedListInt16(t *testing.T) {
	c := is.NewRelaxed(t)
	stack := wypes.NewSliceStack(4)
	store := wypes.Store{Stack: stack, Memory: wypes.NewSliceMemory(1024)}

	data := []wypes.Int16{1, -2, 3, 4, -5, 6, -7, 8, 9, 10}
	wypes.ReturnedList[wypes.Int16]{Offset: 64, Raw: data, DataPtr: 128}.Lower(&store)

	store.Stack.Push(64)
	list := wypes.ReturnedList[wypes.Int16]{}.Lift(&store)

	is.SliceEqual(c, list.Unwrap(), data)
}

func TestReturnedListInt32(t *testing.T) {
	c := is.NewRelaxed(t)
	stack := wypes.NewSliceStack(4)
	store := wypes.Store{Stack: stack, Memory: wypes.NewSliceMemory(1024)}

	data := []wypes.Int32{1, -2, 3, 4, -5, 6, -7, 8, 9, 10}
	wypes.ReturnedList[wypes.Int32]{Offset: 64, Raw: data, DataPtr: 128}.Lower(&store)

	store.Stack.Push(64)
	list := wypes.ReturnedList[wypes.Int32]{}.Lift(&store)

	is.SliceEqual(c, list.Unwrap(), data)
}

func TestReturnedListFloat32(t *testing.T) {
	c := is.NewRelaxed(t)
	stack := wypes.NewSliceStack(4)
	store := wypes.Store{Stack: stack, Memory: wypes.NewSliceMemory(1024)}

	data := []wypes.Float32{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9, 10.1}
	wypes.ReturnedList[wypes.Float32]{Raw: data, DataPtr: 128}.Lower(&store)

	store.Stack.Push(0)
	list := wypes.ReturnedList[wypes.Float32]{}.Lift(&store)

	is.SliceEqual(c, list.Unwrap(), data)
}

func TestListEmpty(t *testing.T) {
	c := is.NewRelaxed(t)
	stack := wypes.NewSliceStack(4)
	store := wypes.Store{Stack: stack, Memory: wypes.NewSliceMemory(1024)}

	wypes.List[wypes.UInt32]{Offset: 64}.Lower(&store)

	store.Stack.Push(64)
	store.Stack.Push(0)
	list := wypes.List[wypes.UInt32]{}.Lift(&store)

	is.SliceEqual(c, list.Unwrap(), []wypes.UInt32{})
}

func TestListUint16(t *testing.T) {
	c := is.NewRelaxed(t)
	stack := wypes.NewSliceStack(4)
	store := wypes.Store{Stack: stack, Memory: wypes.NewSliceMemory(1024)}

	data := []wypes.UInt16{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	wypes.List[wypes.UInt16]{Offset: 96, Raw: data}.Lower(&store)

	store.Stack.Push(96)
	store.Stack.Push(10)
	list := wypes.List[wypes.UInt16]{}.Lift(&store)

	is.SliceEqual(c, list.Unwrap(), data)
}

func TestListUint32(t *testing.T) {
	c := is.NewRelaxed(t)
	stack := wypes.NewSliceStack(4)
	store := wypes.Store{Stack: stack, Memory: wypes.NewSliceMemory(1024)}

	data := []wypes.UInt32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	wypes.List[wypes.UInt32]{Offset: 64, Raw: data}.Lower(&store)

	store.Stack.Push(64)
	store.Stack.Push(10)
	list := wypes.List[wypes.UInt32]{}.Lift(&store)

	is.SliceEqual(c, list.Unwrap(), data)
}

func TestListInt16(t *testing.T) {
	c := is.NewRelaxed(t)
	stack := wypes.NewSliceStack(4)
	store := wypes.Store{Stack: stack, Memory: wypes.NewSliceMemory(1024)}

	data := []wypes.Int16{1, -2, 3, -4, -5, 6, -7, 8, -9, 10}
	wypes.List[wypes.Int16]{Offset: 96, Raw: data}.Lower(&store)

	store.Stack.Push(96)
	store.Stack.Push(10)
	list := wypes.List[wypes.Int16]{}.Lift(&store)

	is.SliceEqual(c, list.Unwrap(), data)
}

func TestListInt32(t *testing.T) {
	c := is.NewRelaxed(t)
	stack := wypes.NewSliceStack(4)
	store := wypes.Store{Stack: stack, Memory: wypes.NewSliceMemory(1024)}

	data := []wypes.Int32{1, -2, 3, -4, -5, 6, -7, 8, -9, 10}
	wypes.List[wypes.Int32]{Offset: 64, Raw: data}.Lower(&store)

	store.Stack.Push(64)
	store.Stack.Push(10)
	list := wypes.List[wypes.Int32]{}.Lift(&store)

	is.SliceEqual(c, list.Unwrap(), data)
}

func TestListFloat32(t *testing.T) {
	c := is.NewRelaxed(t)
	stack := wypes.NewSliceStack(4)
	store := wypes.Store{Stack: stack, Memory: wypes.NewSliceMemory(1024)}

	data := []wypes.Float32{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9, 10.1}
	wypes.List[wypes.Float32]{Raw: data}.Lower(&store)

	store.Stack.Push(0)
	store.Stack.Push(10)
	list := wypes.List[wypes.Float32]{}.Lift(&store)

	is.SliceEqual(c, list.Unwrap(), data)
}

func TestListBool(t *testing.T) {
	c := is.NewRelaxed(t)
	stack := wypes.NewSliceStack(4)
	store := wypes.Store{Stack: stack, Memory: wypes.NewSliceMemory(1024)}

	data := []wypes.Bool{true, false, false, true, true, false, true, false, true, false}
	wypes.List[wypes.Bool]{Offset: 64, Raw: data}.Lower(&store)

	store.Stack.Push(64)
	store.Stack.Push(10)
	list := wypes.List[wypes.Bool]{}.Lift(&store)

	is.SliceEqual(c, list.Unwrap(), data)
}

func TestListEmptyStrings(t *testing.T) {
	c := is.NewRelaxed(t)
	stack := wypes.NewSliceStack(4)
	store := wypes.Store{Stack: stack, Memory: wypes.NewSliceMemory(1024)}

	data := []string{}
	wypes.ListStrings{Offset: 64, Raw: data}.Lower(&store)

	store.Stack.Push(64)
	store.Stack.Push(uint64(len(data)))
	list := wypes.ListStrings{}.Lift(&store)

	is.SliceEqual(c, list.Unwrap(), data)
}

func TestListStrings(t *testing.T) {
	c := is.NewRelaxed(t)
	stack := wypes.NewSliceStack(4)
	store := wypes.Store{Stack: stack, Memory: wypes.NewSliceMemory(1024)}

	data := []string{"Hello", "World", "!"}
	wypes.ListStrings{Offset: 64, Raw: data}.Lower(&store)

	store.Stack.Push(64)
	store.Stack.Push(uint64(len(data)))
	list := wypes.ListStrings{}.Lift(&store)

	is.SliceEqual(c, list.Unwrap(), data)
}

func TestResultOKUInt32(t *testing.T) {
	c := is.NewRelaxed(t)
	stack := wypes.NewSliceStack(4)
	store := wypes.Store{Stack: stack, Memory: wypes.NewSliceMemory(1024)}
	save := wypes.Result[wypes.UInt32, wypes.UInt32, wypes.UInt32]{
		IsError: false,
		OK:      1,
		Offset:  64,
		DataPtr: 128,
	}

	save.Lower(&store)
	store.Stack.Push(64)
	result := wypes.Result[wypes.UInt32, wypes.UInt32, wypes.UInt32]{}.Lift(&store)

	is.Equal(c, result.IsError, false)
	is.Equal(c, result.OK, wypes.UInt32(1))
}

func TestResultErrUInt32(t *testing.T) {
	c := is.NewRelaxed(t)
	stack := wypes.NewSliceStack(4)
	store := wypes.Store{Stack: stack, Memory: wypes.NewSliceMemory(1024)}
	save := wypes.Result[wypes.UInt32, wypes.UInt32, wypes.UInt32]{
		IsError: true,
		Error:   1,
		Offset:  64,
		DataPtr: 128,
	}

	save.Lower(&store)
	store.Stack.Push(64)
	result := wypes.Result[wypes.UInt32, wypes.UInt32, wypes.UInt32]{}.Lift(&store)

	is.Equal(c, result.IsError, true)
	is.Equal(c, result.Error, wypes.UInt32(1))
}

func TestResultOKStringUInt32(t *testing.T) {
	c := is.NewRelaxed(t)
	stack := wypes.NewSliceStack(4)
	store := wypes.Store{Stack: stack, Memory: wypes.NewSliceMemory(1024)}
	save := wypes.Result[wypes.String, wypes.String, wypes.UInt32]{
		IsError: false,
		OK:      wypes.String{Raw: "awesome"},
		Error:   wypes.UInt32(0),
		Offset:  64,
		DataPtr: 128,
	}

	save.Lower(&store)
	store.Stack.Push(64)
	result := wypes.Result[wypes.String, wypes.String, wypes.UInt32]{}.Lift(&store)

	is.Equal(c, result.IsError, false)
	is.Equal(c, result.OK.Unwrap(), "awesome")
}

func TestResultErrStringUInt32(t *testing.T) {
	c := is.NewRelaxed(t)
	stack := wypes.NewSliceStack(4)
	store := wypes.Store{Stack: stack, Memory: wypes.NewSliceMemory(1024)}
	save := wypes.Result[wypes.String, wypes.String, wypes.UInt32]{
		IsError: true,
		OK:      wypes.String{Raw: "awesome"},
		Error:   wypes.UInt32(100),
		Offset:  64,
		DataPtr: 128,
	}

	save.Lower(&store)
	store.Stack.Push(64)
	result := wypes.Result[wypes.String, wypes.String, wypes.UInt32]{}.Lift(&store)

	is.Equal(c, result.IsError, true)
	is.Equal(c, result.Error.Unwrap(), 100)
}

func TestResultOKListUInt32(t *testing.T) {
	c := is.NewRelaxed(t)
	stack := wypes.NewSliceStack(4)
	store := wypes.Store{Stack: stack, Memory: wypes.NewSliceMemory(1024)}
	save := wypes.Result[wypes.List[wypes.UInt32], wypes.List[wypes.UInt32], wypes.UInt32]{
		IsError: false,
		OK:      wypes.List[wypes.UInt32]{Raw: []wypes.UInt32{1, 2, 3, 4, 5}},
		Error:   wypes.UInt32(0),
		Offset:  64,
		DataPtr: 128,
	}

	save.Lower(&store)
	store.Stack.Push(64)
	result := wypes.Result[wypes.List[wypes.UInt32], wypes.List[wypes.UInt32], wypes.UInt32]{}.Lift(&store)

	is.Equal(c, result.IsError, false)
	is.SliceEqual(c, result.OK.Raw, []wypes.UInt32{1, 2, 3, 4, 5})
}

func TestResultErrListUInt32(t *testing.T) {
	c := is.NewRelaxed(t)
	stack := wypes.NewSliceStack(4)
	store := wypes.Store{Stack: stack, Memory: wypes.NewSliceMemory(1024)}
	save := wypes.Result[wypes.List[wypes.UInt32], wypes.List[wypes.UInt32], wypes.String]{
		IsError: true,
		Error:   wypes.String{Raw: "error"},
		Offset:  64,
		DataPtr: 128,
	}

	save.Lower(&store)
	store.Stack.Push(64)
	result := wypes.Result[wypes.List[wypes.UInt32], wypes.List[wypes.UInt32], wypes.String]{}.Lift(&store)

	is.Equal(c, result.IsError, true)
	is.Equal(c, result.Error.Unwrap(), "error")
}

func TestResultOKBytes(t *testing.T) {
	c := is.NewRelaxed(t)
	stack := wypes.NewSliceStack(4)
	store := wypes.Store{Stack: stack, Memory: wypes.NewSliceMemory(1024)}
	save := wypes.Result[wypes.Bytes, wypes.Bytes, wypes.UInt32]{
		IsError: false,
		OK:      wypes.Bytes{Raw: []byte{1, 2, 3, 4, 5}},
		Error:   wypes.UInt32(0),
		Offset:  64,
		DataPtr: 128,
	}

	save.Lower(&store)
	store.Stack.Push(64)
	result := wypes.Result[wypes.Bytes, wypes.Bytes, wypes.UInt32]{}.Lift(&store)

	is.Equal(c, result.IsError, false)
	is.SliceEqual(c, result.OK.Raw, []byte{1, 2, 3, 4, 5})
}
