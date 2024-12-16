package main

/*
#include "print.h"
#include "print.c"
*/
import "C"

func main() {
	C.printPiFromCcode()
}
