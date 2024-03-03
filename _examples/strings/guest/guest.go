package main

import "unsafe"

//go:export print
func hostPrint(string)

//go:wasmimport env get_name
func hostGetName(uint32) (uint32, uint32)

//go:export greet
func greet() {
	b := make([]byte, 20)
	addr := uintptr(unsafe.Pointer(unsafe.SliceData(b)))
	name, _ := hostGetName(uint32(addr))
	_ = name
	// hostPrint("Hello, " + name)
}
