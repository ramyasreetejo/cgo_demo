package main

//Passing Simple Types (int, float, etc.)

/*
#include <stdio.h>

void printInt(int x) {
    printf("Received value: %d\n", x);
}
*/
import "C"

func main() {
	x := 42
	C.printInt(C.int(x)) // Passing int from Go to C
}
