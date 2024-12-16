package main

/*
extern int goAdd(int, int);
static int cAdd(int a, int b){
    return goAdd(a, b);
}
*/
import "C"
import "fmt"

// expose this func as callable from c

//export goAdd
func goAdd(a, b C.int) C.int {
	return a + b
}

// c and go int may not be the same size remember that
func main() {
	fmt.Println(C.cAdd(C.int(1), C.int(3)))
}
