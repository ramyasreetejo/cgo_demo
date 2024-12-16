package main

/*
Go Structs with C Pointers (Cgo Pointers):
If a Go struct contains pointers to C types, passing these objects to C can get complex,
because Go's garbage collector doesn't automatically track memory managed by C.
Therefore, you should manually manage memory and avoid relying on Go's garbage collector to clean up C memory.
*/

/*
#include <stdio.h>
#include <stdlib.h>

struct Data {
    char *name;
};

void printData(struct Data *data) {
    printf("Name: %s\n", data->name);
}
*/
import "C"
import "unsafe"

func main() {
	// Create a Go string
	goStr := "Hello, C!"

	// Convert the Go string to a C string
	cStr := C.CString(goStr)
	defer C.free(unsafe.Pointer(cStr)) // Free the C string after use

	// Create a C struct Data
	data := C.struct_Data{
		name: cStr,
	}

	// Call the C function with the C struct
	C.printData(&data)
}
