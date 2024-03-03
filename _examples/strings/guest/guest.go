package main

//go:export print
func hostPrint(string)

//go:export greet
func greet() {
	hostPrint("Hello, Joe")
}
