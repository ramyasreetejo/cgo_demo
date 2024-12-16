package main

//Passing Cgo Structs to C

/*
#include <stdio.h>

struct Point {
    int x;
    int y;
};

void printPoint(struct Point p) {
    printf("Point: (%d, %d)\n", p.x, p.y);
}
*/
import "C"

type Point struct {
	X int
	Y int
}

func main() {
	p := Point{X: 10, Y: 20}

	// Create a C struct from the Go struct
	// C.struct_Point is the Go representation of the C struct Point.
	cPoint := C.struct_Point{ // Direct mapping to the C struct
		x: C.int(p.X),
		y: C.int(p.Y),
	}

	// Call C function with the C struct
	C.printPoint(cPoint)
}
