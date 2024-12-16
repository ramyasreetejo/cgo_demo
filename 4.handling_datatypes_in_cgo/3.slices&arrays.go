package main

/*
#include <stdio.h>
#include <stdlib.h>

// Function to double the values of the slice
void double_slice(int* slice, int length) {
    for (int i = 0; i < length; i++) {
        slice[i] *= 2;
    }
}
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func main() {
	// Go slice to be passed to C
	goSlice := []int{1, 2, 3, 4, 5}

	// Length of the slice
	l := len(goSlice)

	// Allocate memory for the C array (C.malloc expects a size_t argument)
	cArray := C.malloc(C.size_t(l) * C.size_t(unsafe.Sizeof(C.int(0))))

	// Make sure to free the allocated memory after use
	defer C.free(unsafe.Pointer(cArray))

	// Create a c slice from the C-allocated memory
	cSlice := unsafe.Slice((*C.int)(cArray), l)

	// Copy the elements from the Go slice to the C slice
	for i, v := range goSlice {
		cSlice[i] = C.int(v)
	}

	// Call the C function to modify the C slice (double the values)
	C.double_slice((*C.int)(cArray), C.int(l))

	// Print the modified values from the C slice
	for i := 0; i < l; i++ {
		fmt.Printf("%d ", cSlice[i])
	}
	fmt.Println()
}
