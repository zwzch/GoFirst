package goandc
/*
#include <stdio.h>
#include <stdlib.h>
void hello() {
	printf("Hello, World!\n");
}
char* cstr() {
return "abcde";
}
*/
import "C"
import "fmt"

func test() {
	C.hello()
	cs:=C.cstr()
	fmt.Println(C.GoString(cs))
}



func Mygoandc() {
	//s := "xxx"
	//cs := C.CString(s)
	//fmt.Println(C.display(cs,1,2))
	//C.free(unsafe.Pointer(cs))
	////fmt.Println("20*30=", C.do_test_so_func(20, 30))
	//C.testxxx()
	test()

	}

