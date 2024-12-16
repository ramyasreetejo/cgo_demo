package main

/*
#cgo LDFLAGS: -L. -lexample
#include "libexample.h"
*/
import "C"

func main() {
	// Call the C function
	C.hello_world()
}
