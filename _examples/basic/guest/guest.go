package main

//go:export add_i32
func addI32(int32, int32) int32

//go:export run
func Run() int32 {
	return addI32(9, 5) * 2
}
