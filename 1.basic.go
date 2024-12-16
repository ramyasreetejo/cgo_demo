package main

/*
//this is interpreted as c code
#include <math.h>
#include <stdio.h>
void printPiFromC(){
  printf("Pi: %f\n", M_PI);
}
*/
import "C"
import "fmt"

func main() {
	//printing from c libraries
	fmt.Println("Pi:", C.M_PI)
	C.printPiFromC()
	//call using c package
}
