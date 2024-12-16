package main

//Passing Complex Objects (Pointers to Objects)

/*
#include <stdio.h>
#include <stdlib.h>

struct Point {
    int x;
    int y;
};

void setPoint(struct Point *p, int x, int y) {
    p->x = x;
    p->y = y;
}

void printPoint(struct Point *p) {
    printf("Point: (%d, %d)\n", p->x, p->y);
}
*/
import "C"
import "unsafe"

func main() {
	// Allocate a C struct Point
	p := (*C.struct_Point)(C.malloc(C.sizeof_struct_Point))

	// Ensure the memory is freed after use
	defer C.free(unsafe.Pointer(p))

	// Set values using the C function
	C.setPoint(p, C.int(10), C.int(20))

	// Print values using the C function
	C.printPoint(p)
}
