package main

/*
#cgo CFLAGS: -I../cgo
#cgo LDFLAGS: -L../cgo  -lcgo
#include <stdlib.h>
#include "cgo.h"
*/
import "C"
import (
	"example/hello"
	"example/hello/world"
	"fmt"
	"strings"
	"unsafe"
)

func main() {
	fmt.Println("main")
	hello2.Hello()
	world.World()

	// cgo: pass strings to and from C function
	// char *strncpy2(char *dest, const char *src, size_t n);

	buf := [32]byte{}
	buflen := len(buf)
	str := "hello"
	cs := C.CString(str)
    defer C.free(unsafe.Pointer(cs))

	C.strncpy2(
		(*C.char) (unsafe.Pointer (&buf)),
		cs,
		C.size_t(buflen) - 1)

	fmt.Println("cgo", string(buf[:]))

	// cgo: pass integers to and from C function
	// int *intcpy(int *dest, int src);

	a := 0;
	b := 123;

	C.intcpy(
		(*C.int) (unsafe.Pointer(&a)),
		C.int(b))

	fmt.Println("cgo", a)

	//string search
	string_search()

}

//find sub strings
func string_search() {
	str := "aaaXXaaaYYaaa"
	sub := "aaa"

	for i := 0; i != len(str); i++ {
		j := strings.Index(str[i:], sub)
		if j == -1 {
			break;
		}
		i += j
		fmt.Print(i, " ")
	}
	fmt.Println()
}
