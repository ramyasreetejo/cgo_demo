package main

/*
#include <stdio.h>
#include <stdlib.h>

static void myprint(char* s){
    printf("%s (printed in c) \n", s);
}
*/
import "C"
import "unsafe"

func main() {
	//wrapper cstring function to handle string difference in c and go
	cs := C.CString("hello world")
	//c string allocates memory on c heap
	defer C.free(unsafe.Pointer(cs))
	C.myprint(cs)
}
